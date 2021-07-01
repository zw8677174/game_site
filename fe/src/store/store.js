export default {
    debug: true,
    state: {
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImExNTU1NDYzYzM2MWU3MDM2YTI3NGE4YjQ0ZTI5MTkyIiwicGFzc3dvcmQiOiIyMDJjYjk2MmFjNTkwNzViOTY0YjA3MTUyZDIzNGI3MCIsImV4cCI6MTk4NTE0MDUyMiwiaXNzIjoiZ2luLWJsb2cifQ.iirLTYV8p1utZBn4Rzl5LVPs4D779S1RdDPQgOcMgk4',
        nickName: '陌生人'
    },
    setUserInfo(data) {
        this.state.token = data['token']

    },
    getToken() {
        return this.state.token;
    }

}
