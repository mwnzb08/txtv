import axios from 'axios'
var https = axios
https.defaults.baseURL = '/api'
https.defaults.timeout = 10 * 1000
https.defaults.responseType = 'json'
https.defaults.withCredentials = true
https.interceptors.request.use(request => {
  let i = 0
  let c = ''
  let newUrl = ''
  for (i; i < request.url.length; i++) {
    c = request.url.charAt(i)
    if (c >= 'A' && c <= 'Z') {
      newUrl += '/' + c.toLowerCase()
    } else {
      newUrl += c
    }
  }
  request.url = newUrl
  return request
}, error => {
  return Promise.reject(error)
})
https.interceptors.response.use(response => {
  if (response.status === 200) {
    return Promise.resolve(response)
  } else {
    return Promise.reject(response)
  }
}, error => {
  switch (error.response.status) {
    case 200:
      break
    case 404: window.message.warning('url not found')
      break
    case 408: window.message.error('service error')
      break
    case 500: window.message.error('service error')
      break
    default: window.message.error('service error')
  }
  return Promise.reject(error)
})

export default https
