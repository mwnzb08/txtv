import axios from 'axios'

axios.defaults.baseURL = '/api'
axios.defaults.timeout = 10 * 1000
axios.defaults.responseType = 'json'
axios.defaults.withCredentials = true
// axios.create({
//   withCredentials: true,
//   url: '/api',
//   baseUrl: '127.0.0.1:8082',
//   timeout: 10 * 1000,
//   responseType: 'json',
//   transformRequest: [(request) => {
//     return request
//   }],
//   transformResponse: [(response) => {
//     return response
//   }]
// })
axios.interceptors.request.use((request) => {
  return request
}, error => {
  return Promise.resolve(error)
})
axios.interceptors.response.use((response) => {
  switch (response.status) {
    case 200:
      break
    case 404: window.message.warning('url not found')
      break
    case 408: window.message.error('service error')
      break
    case 500: window.message.error('service error')
      break
    default:
  }
  if (response.status === 200) {
    return Promise.resolve(response)
  } else {
    return Promise.reject(response)
  }
}, error => {
  return Promise.resolve(error)
})

export default axios
