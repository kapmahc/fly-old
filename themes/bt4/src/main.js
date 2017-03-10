import i18next from 'i18next';
import XHR from 'i18next-xhr-backend';
import LanguageDetector from 'i18next-browser-languagedetector';
import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

const locale = 'locale';

function main(){
  i18next
    .use(XHR)
    .use(LanguageDetector)
    .init({
      backend: {
        loadPath: `${process.env.REACT_APP_BACKEND}/locales/{{lng}}`,
        crossDomain: true,
      },
      detection: {
        // order and from where user language should be detected
        order: ['querystring', 'cookie', 'localStorage', 'navigator', 'htmlTag'],
        // keys or params to lookup language from
        lookupQuerystring: locale,
        lookupCookie: locale,
        lookupLocalStorage: locale,
        // cache user language on
        caches: ['localStorage', 'cookie'],
        // optional expire and domain for set cookie
        cookieMinutes: 1<<32-1,
        // optional htmlTag with lang attribute, the default is:
        htmlTag: document.documentElement
      },
    }, (err, t) => {
      ReactDOM.render(
        <App />,
        document.getElementById('root')
      );
  });
}

export default main
