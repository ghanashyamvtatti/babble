import { message } from "antd";
import { FEED, LOGIN, USER } from "../pages";

export const SIGN_IN = "SIGN_IN";
export const SIGN_OUT = "SIGN_OUT";
export const LOAD_FEED = "LOAD_FEED";
export const LOAD_SUBSCRIPTIONS = "LOAD_SUBSCRIPTIONS";
export const LOAD_PROFILE = "LOAD_PROFILE";
export const ADD_POST = "ADD_POST";
export const SIGN_UP = "SIGN_UP";
export const CHANGE_PAGE = "CHANGE_PAGE";
export const LOAD_USERS = "LOAD_USERS";

export const signIn = payload => ({
  type: SIGN_IN,
  payload
});

export const changePage = payload => ({
  type: CHANGE_PAGE,
  payload
});

export function signInProcess(username, password) {
  return async function(dispatch) {
    const res = await fetch("http://localhost:8080/auth/sign-in", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: username,
        password: password
      })
    });
    const res_1 = await res.json();
    console.log(res_1);
    if (res_1.Status) {
      dispatch(signIn(res_1.Data));
      dispatch(changePage(FEED));
    } else {
      message.error(res_1.Message);
    }
  };
}

export const signOut = payload => ({
  type: SIGN_OUT,
  payload
});

export const loadUsers = payload => ({
  type: LOAD_USERS,
  payload
});

export function signOutProcess(username) {
  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/auth/user/" + username + "/sign-out",
      { method: "POST" }
    );
    const res = await resp.json();
    if (res.Status) {
      dispatch(signOut());
      dispatch(changePage(LOGIN));
    }
  };
}

export const loadFeed = payload => ({
  type: LOAD_FEED,
  payload
});

export function loadFeedData(username, token) {
  return async function(dispatch) {
    const res = await fetch(
      "http://localhost:8080/social/user/" + username + "/feed",
      {
        method: "GET",
        headers: {
          token: token
        }
      }
    );
    const data = await res.json();
    if (data.Status) {
      if (data.Data.feed === undefined || data.Data.feed == null) {
        dispatch(loadFeed([]));
      } else {
        dispatch(loadFeed(data.Data.feed));
      }
    } else {
      message.error(data.Message);
    }
  };
}

export const loadSubscriptions = payload => ({
  type: LOAD_SUBSCRIPTIONS,
  payload
});

export function loadSubscriptionDetails(token, username) {
  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/social/user/" + username + "/subscriptions",
      {
        method: "GET",
        headers: {
          token: token
        }
      }
    );
    const data = await resp.json();

    if (data.Status) {
      if (
        data.Data.subscriptions === undefined ||
        data.Data.subscriptions == null
      ) {
        data.Data.subscriptions = [];
        dispatch(loadSubscriptions(data.Data));
      } else {
        dispatch(loadSubscriptions(data.Data));
      }
    } else {
      message.error(data.Message);
    }
  };
}

export function unsubscribe(token, username, publisher) {
  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/social/user/" +
        username +
        "/subscribe/" +
        publisher,
      {
        method: "DELETE",
        headers: {
          token: token
        }
      }
    );
    const data = await resp.json();
    if (data.Status) {
      dispatch(loadSubscriptionDetails(token, username));
      message.success("Successfully unsubscribed");
    } else {
      message.error(data.Message);
    }
  };
}

export function subscribe(token, username, publisher) {
  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/social/user/" +
        username +
        "/subscribe/" +
        publisher,
      {
        method: "POST",
        headers: {
          token: token
        }
      }
    );
    const data = await resp.json();
    if (data.Status) {
      dispatch(loadSubscriptionDetails(token, username));
      message.success("Successfully subscribed");
    } else {
      message.error(data.Message);
    }
  };
}

export function getAllUsers(token) {
  return async function(dispatch) {
    const resp = await fetch("http://localhost:8080/social/user");
    const data = await resp.json();
    if (data.Status) {
      dispatch(loadUsers(data.Data.result));
    } else {
      message.error(data.Message);
    }
  };
}

export const loadUserProfile = payload => ({
  type: LOAD_PROFILE,
  payload
});

export function loadUserDetails(token, me, username) {
  console.log(token, me, username);

  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/social/user/" + me + "/?username=" + username,
      {
        method: "GET",
        headers: {
          token: token
        }
      }
    );
    const res = await resp.json();
    if (res.Status) {
      console.log(res);
      var user = res.Data.user;
      const postResp = await fetch(
        "http://localhost:8080/social/user/" +
          me +
          "/post?username=" +
          username,
        {
          headers: {
            token: token
          }
        }
      );
      const postRes = await postResp.json();
      if (postRes.Status) {
        var posts = postRes.Data.posts;
        if (posts === undefined || posts == null) {
          dispatch(loadUserProfile({ user: user, posts: [] }));
        } else {
          dispatch(loadUserProfile({ user: user, posts: posts }));
        }
        dispatch(changePage(USER));
      } else {
        message.error(postRes.Message);
      }
    } else {
      message.error(res.Message);
    }
  };
}

export const addPost = payload => ({
  type: ADD_POST,
  payload
});

export function addPostProcess(token, username, postText) {
  console.log(token, username, postText);
  return async function(dispatch) {
    const resp = await fetch(
      "http://localhost:8080/social/user/" + username + "/post",
      {
        method: "POST",
        headers: {
          token: token,
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ post: postText })
      }
    );
    const res = await resp.json();
    if (res.Status) {
      console.log(res);
      dispatch(loadFeedData(username, token));
    } else {
      message.error(res.Message);
    }
  };
}

export const signUp = payload => ({
  type: SIGN_UP,
  payload
});

export function signUpProcess(name, username, password) {
  console.log(username, password, name);
  return async function(dispatch) {
    const res = await fetch("http://localhost:8080/auth/sign-up", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        name: name,
        username: username,
        password: password
      })
    });
    const res_1 = await res.json();
    if (res_1.Status) {
      console.log(res_1);
      dispatch(signUp(res_1.Data));
      dispatch(changePage(FEED));
    } else {
      message.error(res_1.Message);
    }
  };
}
