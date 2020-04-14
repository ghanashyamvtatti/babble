import React, { Component } from "react";
import { connect } from "react-redux";

import { loadFeedData } from "../../actions/actions";
import Post from "../Post/Post";
import AddPost from "../AddPost/AddPost";

class Feed extends Component {
  componentDidMount() {
    console.log("Calling loadFeedData");

    this.props.loadFeedData(this.props.username, this.props.token);
  }

  render() {
    return (
      <div>
        {this.props.posts.map(function(post, index) {
          return (
            <Post
              key={index}
              Username={post.username}
              Post={post.post}
              CreatedAt={post.created_at.seconds}
            />
          );
        })}
        <AddPost />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  token: state.token,
  username: state.me.username,
  posts: state.posts
});

export default connect(mapStateToProps, { loadFeedData })(Feed);
