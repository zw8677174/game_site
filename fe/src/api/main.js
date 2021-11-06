const axios = require('axios');

// axios.interceptors.response.use((response) => {
//     return response.data.data;
// })

import store from "@/store/store.js"

export default {

    name: 'api',
    version: 'v1.0',
    initParams: function (params) {
        params.token = store.getToken();
        return params;
    },

    // auth
    auth: function (data, func) {
        return axios.post('/api/auth', data).then(v => func(v.data.data));
    },

    getGameList: function(data = {}) {
        return axios({
            url: '/api/v1/game_list',
            method: 'GET',
            params: this.initParams(data)
        });
        // return axios.get('/api/v1/game_list', this.initParams(data)).then(v => func(v.data.data));
    },


}

