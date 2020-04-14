import { Avatar, Card, Comment, Tooltip } from "antd";
import React, { Component } from "react";
import { loadUserDetails } from "../../actions/actions.js";
import { connect } from "react-redux";

class Post extends Component {
  goToUserPage = () => {
    this.props.loadUserDetails(
      this.props.token,
      this.props.me,
      this.props.Username
    );
  };
  render() {
    return (
      <Card
        style={{
          maxWidth: "50%",
          background: "#177ddc",
          marginLeft: "25%",
          marginRight: "25%",
          marginBottom: "25px"
        }}
        bordered={false}
        hoverable={true}
        loading={false}
      >
        <Comment
          // eslint-disable-next-line jsx-a11y/anchor-is-valid
          author={<a>{this.props.Username}</a>}
          avatar={
            <Avatar alt={this.props.Username} onClick={this.goToUserPage}>
              {this.props.Username[0].toUpperCase()}
            </Avatar>
          }
          content={<p>{this.props.Post}</p>}
          datetime={
            <Tooltip
              title={new Date(this.props.CreatedAt * 1000).toLocaleString()}
            >
              <span>
                {new Date(this.props.CreatedAt * 1000).toLocaleString()}
              </span>
            </Tooltip>
          }
        />
      </Card>
    );
  }
}
const mapStateToProps = state => ({
  token: state.token,
  me: state.me.username
});
export default connect(mapStateToProps, { loadUserDetails })(Post);
