const apiHost = 'http://localhost:3000'
export const currentLocale = () => localStorage.getItem('locale') || 'zh-Hans'

export const loadLocaleMessage = (locale, cb) => {
  return fetch(`${apiHost}/locales/${locale}`, {
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
    cb(null, message)
  }).catch((error) => {
    cb(error)
  })
}
