import ApiService from './api.service'

const HomeService = {
  fetchPosts() {
    console.log('fetch')
    return ApiService.get('/posts')
  },
}

export default HomeService
