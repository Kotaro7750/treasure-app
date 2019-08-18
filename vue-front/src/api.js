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

export const createArticle = function (idToken, title, body, tags, jiro) {
    return fetch(`${API_ENDPOINT}/articles`, {
        method: "POST",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin",
        body: JSON.stringify({ article: { title: title, body: body }, tags: tags, jiro: jiro }),
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

export const deleteArticle = function (idToken, article_id) {
    return fetch(`${API_ENDPOINT}/articles/` + article_id, {
        method: "DELETE",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin",
    }).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`Request rejected with status ${res.status}`);
        }
    });
}

export const getTagList = function () {
    return fetch(`${API_ENDPOINT}/tags`).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`${res.status}`)
        }
    })
};


export const createComment = function (idToken, article_id, body) {
    return fetch(`${API_ENDPOINT}/articles/` + article_id + `/comments`, {
        method: "POST",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin",
        body: JSON.stringify({ body: body }),
    }).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`Request rejected with status ${res.status}`);
        }
    });
};

//jiro
export const showJiro = function (jiro_id) {
    return fetch(`${API_ENDPOINT}/jiros/` + jiro_id).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`${res.status}`)
        }
    })
};
export const getJiroList = function () {
    return fetch(`${API_ENDPOINT}/jiros`).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`${res.status}`)
        }
    })
};

export const nearestJiro = function (idToken, position) {
    return fetch(`${API_ENDPOINT}/jiros/nearest`, {
        method: "POST",
        headers: new Headers({
            Authorization: `Bearer ${idToken}`
        }),
        credentials: "same-origin",
        body: JSON.stringify({ position: position })
    }).then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error(`Request rejected with status ${res.status}`);
        }
    });
};