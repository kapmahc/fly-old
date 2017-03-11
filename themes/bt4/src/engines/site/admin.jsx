import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {ListGroup, ListGroupItem, FormGroup,ControlLabel, FormControl, Button} from 'react-bootstrap';
import i18next from 'i18next';

import {get, post} from '../../ajax'


class InfoW extends Component{
  constructor(props){
    super(props)
    this.state = {
      title:'',
      subTitle:'',
      keywords:'',
      description: '',
      copyright: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentWillReceiveProps(nextProps) {
    const {info} = nextProps
    this.setState(info)
  }
  componentDidMount(){
    const {info} = this.props
    if(info){
      this.setState(info)
    }
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('title', this.state.title)
    data.append('subTitle', this.state.subTitle)
    data.append('keywords', this.state.keywords)
    data.append('description', this.state.description)
    data.append('copyright', this.state.copyright)
    post('/admin/site/info', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.info.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="title">
          <ControlLabel>{i18next.t('site.attributes.title')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.title}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="subTitle">
          <ControlLabel>{i18next.t('site.attributes.subTitle')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.subTitle}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="keywords">
          <ControlLabel>{i18next.t('site.attributes.keywords')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.keywords}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="description">
          <ControlLabel>{i18next.t('site.attributes.description')}</ControlLabel>
          <FormControl
            value={this.state.description}
            onChange={this.handleChange}
            componentClass="textarea" />
        </FormGroup>
        <FormGroup controlId="copyright">
          <ControlLabel>{i18next.t('site.attributes.copyright')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.copyright}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}

InfoW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Info = connect(
  state => ({info:state.siteInfo})
)(InfoW);

class AuthorW extends Component{
  constructor(props){
    super(props)
    this.state = {
      name:'',
      email: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentWillReceiveProps(nextProps) {
    const {info} = nextProps
    this.setState(info.author)
  }
  componentDidMount(){
    const {info} = this.props
    if(info.author){
      this.setState(info.author)
    }
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('name', this.state.name)
    data.append('email', this.state.email)
    post('/admin/site/author', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.author.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="name">
          <ControlLabel>{i18next.t('attributes.fullName')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.name}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="email">
          <ControlLabel>{i18next.t('attributes.email')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.email}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}

AuthorW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Author = connect(
  state => ({info:state.siteInfo})
)(AuthorW);

class SeoW extends Component{
  constructor(props){
    super(props)
    this.state = {
      googleVerifyCode:'',
      baiduVerifyCode: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/admin/site/seo').then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('googleVerifyCode', this.state.googleVerifyCode)
    data.append('baiduVerifyCode', this.state.baiduVerifyCode)
    post('/admin/site/seo', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    const {info} = this.props
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.seo.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="googleVerifyCode">
          <ControlLabel>{i18next.t('site.attributes.googleVerifyCode')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.googleVerifyCode}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="baiduVerifyCode">
          <ControlLabel>{i18next.t('site.attributes.baiduVerifyCode')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.baiduVerifyCode}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
      <br/>
      <ListGroup>
        {info.languages.map((l)=>(`rss-${l}.atom`)).concat(["sitemap.xml.gz", "robots.txt", `google${this.state.googleVerifyCode}.html`, `baidu_verify_${this.state.baiduVerifyCode}.html`]).map((f,i)=>(
          <ListGroupItem key={i} target="_blank" href={`/${f}`}>{f}</ListGroupItem>
        ))}
      </ListGroup>
    </div>)
  }
}

SeoW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Seo = connect(
  state => ({info:state.siteInfo})
)(SeoW);

export class Smtp extends Component{
  constructor(props){
    super(props)
    this.state = {
      name:'',
      email: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/site/info').then(
      function(rst){
        this.setState(rst.author)
      }.bind(this)
    );
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('name', this.state.name)
    data.append('email', this.state.email)
    post('/admin/site/author', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.author.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="name">
          <ControlLabel>{i18next.t('attributes.fullName')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.name}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="email">
          <ControlLabel>{i18next.t('attributes.email')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.email}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}
