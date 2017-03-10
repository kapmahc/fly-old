import React from 'react';

import Header from './components/Header'
import Footer from './components/Footer'

export const Widget = ({children}) => (
  <div>
    <Header />
    <div className="container">
      {children}
    </div>
    <Footer/>
  </div>
)

export default Widget;
