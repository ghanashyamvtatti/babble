import React, { Component } from "react";
import { connect } from "react-redux";
import { Typography } from "antd";
import { loadFeedData } from "../../actions/actions";
import Post from "../Post/Post";

const { Title } = Typography;

class UserPage extends Component {
  render() {
    return (
      <div>
        <Title>{this.props.user.username}</Title>
        {this.props.posts.map(function(post, index) {
          return (
            <Post
              key={index}
              Username={post.username}
              Post={post.post}
              CreatedAt={post.created_at}
            />
          );
        })}
      </div>
    );
  }
}

const mapStateToProps = state => ({
  token: state.token,
  username: state.me.username,
  user: state.user,
  posts: state.posts
});

export default connect(mapStateToProps, { loadFeedData })(UserPage);
