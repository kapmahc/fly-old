import Vue from 'vue'
import I18n from 'vue-i18n'
import Cookies from 'js-cookie'

Vue.use(I18n)

import {api} from '@/ajax'
import {LOCALE} from '@/constants'

const locale = localStorage.getItem(LOCALE) || process.env.LOCALE
Cookies.set(LOCALE, locale, { expires: 7 })
const messages = {}
messages[locale] = {}
const i18n = new I18n({ locale, messages })

const loadLocaleMessage = (l) => {
  return fetch(api(`/locales/${l}`), {
    method: 'get',
    mode: 'cors',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  }).then((res) => {
    return res.json()
  }).then((json) => {
    if (Object.keys(json).length === 0) {
      return Promise.reject(new Error('locale empty !!'))
    } else {
      return Promise.resolve(json)
    }
  }).then((message) => {
    i18n.setLocaleMessage(l, message)
  }).catch((error) => {
    console.log(error)
  })
}

loadLocaleMessage(locale)

export default i18n
