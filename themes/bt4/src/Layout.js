import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'

import {refresh} from './actions'
import {get} from './ajax'

class Widget extends Component {
  componentDidMount() {
    const { refresh } = this.props
    get('/site/info').then(
      rst => {
        document.title = rst.title;
        refresh(rst);
      }
    );
  }
  render() {
    const {children} = this.props;
    return (
      <div>
        <Header />
        <div className="container">
          {children}
        </div>
        <Footer/>
      </div>
    );
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  refresh: PropTypes.func.isRequired
}

export default connect(
  state => ({}),
  {refresh},
)(Widget);
