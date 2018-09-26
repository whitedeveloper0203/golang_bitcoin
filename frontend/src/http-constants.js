import axios from 'axios'

let baseURL

if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
  // where our go API is served at. start from here when calling the API
  baseURL = 'http://127.0.0.1:8000/'
} else {
  baseURL = 'http://api.example.com'
}

export const HTTP = axios.create({
  baseURL: baseURL
})

/*
and then in every .vue file needing Axios, import HTTP instead of Axios:

import {HTTP} from '../http-constants'

HTTP.get(...).then(...).catch(...)

*/
