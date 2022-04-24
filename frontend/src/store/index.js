import {createStore} from 'vuex'
import http from '../http/index'

const loaclToken = localStorage.getItem("token")

// 请求服务器接口获取用户信息。
// 将用户信息存到vuex中

export default createStore({
    state:{
        tokenInfo: null,
    },
    mutations:{
        getTokenInfo(state,value){
            state.tokenInfo = value
        }
    },
    actions:{
        // setTokenInfo(context,value){
        //     context.commit("getTokenInfo",value)
        // },
        getToken(context){
            return new Promise((reslove,reject)=>{
                http.post("/getTokenInfo").then(res=>{
                    context.commit("getTokenInfo",res.data.data)
                    reslove(res)
                }).catch(res=>{reject(res)})
            })
        }
    }
})

