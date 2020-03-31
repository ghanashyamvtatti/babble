import {
  SIGN_IN,
  SIGN_OUT,
  LOAD_FEED,
  LOAD_SUBSCRIPTIONS,
  LOAD_PROFILE,
  ADD_POST,
  SIGN_UP,
  CHANGE_PAGE,
  LOAD_USERS
} from "../actions/actions";
import { LOGIN } from "../pages";

const initialState = {
  page: LOGIN,
  token: "",
  me: {},
  user: {},
  posts: [],
  subscriptions: [],
  data: []
};

export default (state = initialState, { type, payload }) => {
  switch (type) {
    case SIGN_UP:
      return { ...state, me: payload.user, token: payload.token };
    case SIGN_IN:
      return { ...state, me: payload.user, token: payload.token };
    case SIGN_OUT:
      return initialState;
    case LOAD_FEED:
      var x = { ...state, posts: payload };
      console.log(x);

      return x;
    case LOAD_SUBSCRIPTIONS:
      return { ...state, subscriptions: payload.subscriptions };
    case LOAD_PROFILE:
      return { ...state, user: payload.user, posts: payload.posts };
    case ADD_POST:
      return state;
    case CHANGE_PAGE:
      return { ...state, page: payload };
    case LOAD_USERS:
      return { ...state, data: payload };
    default:
      return state;
  }
};
