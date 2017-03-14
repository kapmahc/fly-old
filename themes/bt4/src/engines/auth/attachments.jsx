import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {Table,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import Dropzone from 'react-dropzone'

import {get, post, _delete} from '../../ajax'
import {Forbidden} from '../../components/alerts'



export class Uploader extends Component{
  constructor(props){
    super(props)
    this.state = {}
    this.handleDrop = this.handleDrop.bind(this);
  }
  componentDidMount() {
  }
  handleDrop(files) {
    var data = new FormData()
    files.forEach((file)=> {      
      data.append('files', file)
    });
    post('/attachments', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<Dropzone onDrop={this.handleDrop}>
      <div>{i18next.t('hints.uploader')}</div>
    </Dropzone>)
  }
}

class DashboardW extends Component{
  constructor(props){
    super(props)
    this.state = {
      items:[]
    }
  }
  componentDidMount() {
    get(`/attachments`).then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  render() {
    const {user} = this.props
    return user.uid ? (<div className="row">
      <h3>{i18next.t('auth.attachments.index.title')}</h3>
      <hr/>
      <Uploader/>
      <br/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>
              {i18next.t('attributes.name')}
            </th>
            <th>
              {i18next.t('buttons.manage')}
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((a,i)=>(<tr key={i}>
            <td>{a.updatedAt}</td>
            <td><a href={a.url} target="_blank">{a.title}</a></td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <Button bsStyle="warning" onClick={()=>this.handleEdit(a)}>{i18next.t('buttons.edit')}</Button>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(a.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
    </div>):<Forbidden/>
  }
}

DashboardW.propTypes = {
  user: PropTypes.object.isRequired
}

export const Dashboard = connect(
  state => ({user:state.currentUser})
)(DashboardW);
