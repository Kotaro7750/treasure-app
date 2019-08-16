const API_ENDPOINT = process.env.VUE_APP_BACKEND_API_BASE;

//auth
export const getPrivateMessage = function (idToken) {
    return fetch(`${API_ENDPOINT}/private`, {
        method: "get",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin"
    }).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`Request rejected with status ${res.status}`);
        }
    });
};

export const getPublicMessage = function () {
    return fetch(`${API_ENDPOINT}/public`);
};

//article
export const showArticle = function (article_id) {
    return fetch(`${API_ENDPOINT}/articles/` + article_id).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`${res.status}`)
        }
    })
};

export const createArticle = function (idToken, title, body) {
    return fetch(`${API_ENDPOINT}/articles`, {
        method: "POST",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin",
        body: JSON.stringify({ title, body }),
    }).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`Request rejected with status ${res.status}`);
        }
    });
};

export const getArticleList = function () {
    return fetch(`${API_ENDPOINT}/articles`).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`${res.status}`)
        }
    })
};
