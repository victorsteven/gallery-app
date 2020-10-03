let API_ROUTE 

process.env.NODE_ENV == 'development'
  ? API_ROUTE = 'http://127.0.0.1:7070'
  : API_ROUTE = ''

export default API_ROUTE