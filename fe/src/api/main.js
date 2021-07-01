const axios = require('axios');

import store from "@/store/store.js"

export default {

    name: 'api',
    version: 'v1.0',
    initAuth: function (params) {
        params['token'] = store.getToken();
        return params;
    },

    // auth
    auth: function (data, func) {
        return axios.post('/api/auth', data).then(v => func(v.data.data));
    },

    getGameList: function(data = [], func) {
        return axios.get('/api/v1/game_list', this.initAuth(data)).then(v => func(v.data.data));
    },


}

