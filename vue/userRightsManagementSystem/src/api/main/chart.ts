import hyRequest from '..'

export function chartData(url: string) {
  return hyRequest.get({ url: url })
}
