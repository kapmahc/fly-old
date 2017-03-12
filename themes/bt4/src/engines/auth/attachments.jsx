import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {Table, FormGroup,ControlLabel, FormControl, Thumbnail,
  HelpBlock,Pagination,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'
import {browserHistory, Link} from 'react-router'

import {get, post, _delete} from '../../ajax'
import Markdown from '../../components/Markdown'
import {PAGE_SIZE} from '../../constants'
import {Forbidden} from '../../components/alerts'

class DashboardW extends Component{
  constructor(props){
    super(props)
    this.state = {
      page:1,
      size:PAGE_SIZE,
      total:0,
      count:1,
      items:[]
    }
    this.handleSelect = this.handleSelect.bind(this);
  }
  componentDidMount() {
    this.loadArticles(1, PAGE_SIZE)
  }
  handleSelect(page){
    this.loadArticles(page, PAGE_SIZE)
  }
  loadArticles(page, size){
    get(`/forum/articles?page=${page}&size=${size}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render() {
    const {user} = this.props
    var pager = (<div className="col-md-12">
      <Pagination
          prev
          next
          first
          last
          ellipsis
          boundaryLinks
          items={this.state.count}
          maxButtons={5}
          activePage={this.state.page}
          onSelect={this.handleSelect} />
      </div>)
    return user.uid ? (<div className="row">
      <h3>{i18next.t('auth.attachments.index.title')}</h3>
      <hr/>
      {pager}

      {pager}
    </div>):<Forbidden/>
  }
}

DashboardW.propTypes = {
  user: PropTypes.object.isRequired
}

export const Dashboard = connect(
  state => ({user:state.currentUser})
)(DashboardW);
