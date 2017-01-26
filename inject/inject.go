package inject

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// Object An Object in the Graph.
type Object struct {
	Value        interface{}
	Name         string
	reflectType  reflect.Type
	reflectValue reflect.Value
}

func (p *Object) isStructPtr() bool {
	return p.reflectType.Kind() == reflect.Ptr && p.reflectType.Elem().Kind() == reflect.Struct
}

// New new injector
func New() *Injector {
	return &Injector{args: make([]*Object, 0)}
}

// Injector represents an interface for mapping and injecting dependencies into structsand function arguments.
type Injector struct {
	args  []*Object
	debug bool
}

// Debug set debug mode
func (p *Injector) Debug(b bool) {
	p.debug = b
}

// Map add a object to graph
func (p *Injector) Map(args ...interface{}) {
	for _, v := range args {
		p.MapTo("", v)
	}
}

// MapTo add a object to graph by name
func (p *Injector) MapTo(n string, v interface{}) {
	p.args = append(
		p.args, &Object{
			Name:         n,
			Value:        v,
			reflectType:  reflect.TypeOf(v),
			reflectValue: reflect.ValueOf(v),
		},
	)
}

// Populate populate the incomplete Objects
func (p *Injector) Populate() error {
	for _, o := range p.args {
		// check type
		if o.Name == "" && !o.isStructPtr() {
			return fmt.Errorf(
				"expected unnamed object value to be a pointer to a struct but got type %s with value %v",
				o.reflectType,
				o.Value,
			)
		}
		// check twice
		var found *Object
		for _, a := range p.args {
			if o.reflectType.AssignableTo(a.reflectType) && o.Name == a.Name {
				if found != nil {
					return fmt.Errorf(
						"found two assignable values in type %s with name %s. one type %s with value %v and another type %s with value %v",
						o.reflectType,
						o.Name,
						found.reflectType,
						found.Value,
						a.reflectType,
						a.reflectValue,
					)
				}
			}
		}

		if o.isStructPtr() {
			if err := p.populateExplicit(o); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Injector) debugf(format string, args ...interface{}) {
	if p.debug {
		log.Printf(format, args...)
	}
}

func (p *Injector) populateExplicit(o *Object) error {
	for i := 0; i < o.reflectValue.Elem().NumField(); i++ {
		field := o.reflectValue.Elem().Field(i)
		fieldType := field.Type()
		fieldTag := o.reflectType.Elem().Field(i).Tag
		fieldName := o.reflectType.Elem().Field(i).Name

		if !field.CanSet() {
			return fmt.Errorf(
				"inject requested on unexported field %s in type %s",
				o.reflectType.Elem().Field(i).Name,
				o.reflectType,
			)
		}

		// Don't overwrite existing values.
		if !p.isNilOrZero(field, fieldType) {
			continue
		}
		// Don't overwrite without tag values.
		name, ok := fieldTag.Lookup("inject")
		if !ok {
			continue
		}

		found := p.get(name, fieldType)

		if found == nil {
			if name != "" {
				return fmt.Errorf("unhandled named instance type %s with name %s", fieldType, name)
			}
			// create default struct
			newValue := reflect.New(fieldType.Elem())
			p.debugf(
				"assigned newly created %+v to field %s in %+v",
				newValue, fieldName, o,
			)
			p.Map(newValue.Interface())
			field.Set(newValue)
			if err := p.Populate(); err != nil {
				return err
			}
		} else {
			p.debugf(
				`assigned existing "%+v" to interface field %s in %+v`,
				found.Value, fieldName, o.Value,
			)
			field.Set(reflect.ValueOf(found.Value))
		}
	}
	return nil
}

func (p *Injector) isNilOrZero(v reflect.Value, t reflect.Type) bool {
	switch v.Kind() {

	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(t).Interface())
	}
}

// Walk walk objects
func (p *Injector) Walk(fn func(*Object) error) error {
	for _, o := range p.args {
		if err := fn(o); err != nil {
			return err
		}
	}
	return nil
}

// Invoke attempts to call the interface{} provided as a function
func (p *Injector) Invoke(f interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("must func")
	}

	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		aty := t.In(i)
		val := p.get("", aty)
		if val == nil {
			return nil, fmt.Errorf("value not found for type %v", aty)
		}

		in[i] = val.reflectValue
		p.debugf("put value %+v to type %s", val.Value, aty)
	}

	return reflect.ValueOf(f).Call(in), nil
}

func (p *Injector) get(n string, t reflect.Type) *Object {

	for _, o := range p.args {
		if o.reflectType.AssignableTo(t) && n == o.Name {
			return o
		}
	}
	return nil
}
