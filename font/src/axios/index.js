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
    case 200: break
    case 404: break
    case 408: break
    case 500: alert('service error')
      break
    default:
  }
}, error => {
  return Promise.resolve(error)
})

export default axios
