const axiosBase = require("axios")

class nana {
    constructor() {
        this.axios = axiosBase.create({
            baseURL: 'https://jackson.nana-music.com/v2',
            headers: {
                'X-Ios-Device': 'iPad4,1',
                'X-Ios-Version': '10.3.2',
                'Accept-Language': 'ja-US;q=1',
                'Content-Type': 'application/json',
                'User-Agent': 'nana/2.24.2 (iPad; iOS 10.3.2; Scale/2.00)',
                'X-App-Version': '2.24.2'
            }
        });
        this.logined = false;
    }

    async accountCreate(name) {
        let deviceid = this.deviceidCreate();
        let result = await this.axios.post('/signup', { screen_name: name, device_id: deviceid });

        return { user: result.data.data.user, token: result.data.data.token, deviceid: deviceid };

    }

    async loginToken(token, deviceid) {
        this.axios = axiosBase.create({
            baseURL: 'https://jackson.nana-music.com/v2',
            headers: {
                'X-Ios-Device': 'iPad4,1',
                'X-Ios-Version': '10.3.2',
                'Accept-Language': 'ja-US;q=1',
                'Content-Type': 'application/json',
                'User-Agent': 'nana/2.24.2 (iPad; iOS 10.3.2; Scale/2.00)',
                'X-App-Version': '2.24.2',
                'Authorization': `Token ${token}`,
                'X-Device-Identifier': deviceid
            }
        });
        try {
            let result = await this.axios.get('/my_page/oauth');

            this.logined = true;
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }

    }
    async musicClap(id) {
        if (!this.logined) {
            throw ('Invalid token')
        }
        try {
            let result = await this.axios.post(`/posts/${id}/applause`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }

    async musicPlay(id) {
        if (!this.logined) {
            throw ('Invalid token')
        }
        try {
            let result = await this.axios.post(`/posts/${id}/play`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }


    async musicInfo(id) {
        try {
            let result = await this.axios.get(`/posts/${id}`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }

    async userProfile(id) {
        try {
            let result = await this.axios.get(`/users/${id}`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }

    async userFollow(id) {
        if (!this.logined) {
            throw ('Invalid token')
        }
        try {
            let result = await this.axios.post(`/users/${id}/follow`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }
    async userContents(id) {
        try {
            let result = await this.axios.get(`/users/${id}/contents`);
            return result.data;
        } catch (e) {
            console.log(e)
            throw (e.response.data.detail);
        }
    }

    deviceidCreate() {
        return `${this.randomHex(8)}-${this.randomHex(4)}-${this.randomHex(4)}-${this.randomHex(4)}-${this.randomHex(12)}`;
    }

    randomHex(len) {
        let maxlen = 8;
        let min = Math.pow(16, Math.min(len, maxlen) - 1);
        let max = Math.pow(16, Math.min(len, maxlen)) - 1;
        let n = Math.floor(Math.random() * (max - min + 1)) + min;
        let r = n.toString(16);
        while (r.length < len) {
            r = r + this.randomHex(len - maxlen);
        }
        return r;
    };

}
module.exports = nana
