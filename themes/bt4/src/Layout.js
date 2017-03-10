import React, { Component, PropTypes } from 'react';
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'

import {refresh} from './actions'

class Widget extends Component {
  componentDidMount() {
    const { refresh } = this.props
    refresh({title:'ttt', subTitle:'sss'});
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
