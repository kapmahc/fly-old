import enUS from './en-US'
import zhHans from './zh-Hans'
import zhHant from './zh-Hant'

export default {
  locale: localStorage.getItem('locale') || 'zh-Hant',
  messages: {
    'en-US': enUS,
    'zh-Hans': zhHans,
    'zh-Hant': zhHant
  }
}
