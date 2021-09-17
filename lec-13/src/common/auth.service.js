import ApiService from './api.service'

const AuthService = {
  register(credentials) {
    return ApiService.post('users', { user: credentials })
  },
}

export default AuthService
