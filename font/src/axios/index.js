import axios from 'axios'

axios.create({
  withCredentials: true,
  url: '/api/',
  baseUrl: '',
  timeout: 10 * 1000,
  responseType: 'json',
  transformRequest: [(request) => {
    if (request) {
      return request
    }
  }],
  transformResponse: [(response) => {
    return response
  }]
})

axios.interceptors.request.use()
axios.interceptor.response.user()

export default axios
