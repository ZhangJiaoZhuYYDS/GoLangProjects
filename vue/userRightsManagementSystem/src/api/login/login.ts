import hyRequest from '..'

export function getCode() {
  return hyRequest.get({
    url: '/auth/code'
  })
}

export function login(requestUser: object) {
  return hyRequest.post({
    url: '/auth/login',
    data: requestUser
  })
}

export function loginToken(token: string) {
  return hyRequest.post({
    url: '/auth/loginByToken?token=' + token
  })
}
