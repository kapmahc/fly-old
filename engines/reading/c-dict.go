package reading

import gin "gopkg.in/gin-gonic/gin.v1"

type fmDict struct {
	Keywords string `form:"keywords" binding:"required,max=255"`
}

func (p *Engine) postDict(c *gin.Context) (interface{}, error) {

	var fm fmDict
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	rst := gin.H{}
	for _, dic := range dictionaries {
		for _, sen := range dic.Translate(fm.Keywords) {
			rst[dic.GetBookName()] = sen.Parts
		}
	}

	return rst, nil
}
