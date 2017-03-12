import React, {Component } from 'react'

export class List extends Component{
  constructor(props){
    super(props)
    this.state = {      
    }
  }
  render() {
    return (<div className="row">
      articles list
    </div>)
  }
}

export class Index extends Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }
  render() {
    return (<div className="row">
      articles index
    </div>)
  }
}

export class New extends Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }
  render() {
    return (<div className="row">
      articles new
    </div>)
  }
}

export class Show extends Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }
  render() {
    return (<div className="row">
      article show
    </div>)
  }
}

export class Edit extends Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }
  render() {
    return (<div className="row">
      articles edit
    </div>)
  }
}

export class Dashboard extends Component{
  constructor(props){
    super(props)
    this.state = {
    }
  }
  render() {
    return (<div className="row">
      articles ds
    </div>)
  }
}
