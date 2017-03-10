import React,{PropTypes} from 'react';
import { connect } from 'react-redux'
import {Link} from 'react-router'
import i18next from 'i18next';

const Widget = ({info}) => (
  <div className="container">
    <hr/>
    <footer>
      <p className="pull-right">
        {i18next.t('footer.other-languages')}
        {info.languages.map((lng,i)=>(
          <Link className="block" to={{ pathname: '/', query: { locale: lng } }} target="_blank">
            {i18next.t(`languages.${lng}`)}
        </Link>
        ))}
      </p>
      <p>
        &copy; {info.copyright}
        &middot; <a href="#">Privacy</a>
        &middot; <a href="#">Terms</a>
      </p>
    </footer>
  </div>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({info: state.siteInfo})
)(Widget)
