import { Avatar, Card, Comment, Tooltip } from "antd";
import React, { Component } from "react";
import moment from "moment";

class Post extends Component {
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
            <Avatar alt={this.props.Username}>
              {this.props.Username[0].toUpperCase()}
            </Avatar>
          }
          content={<p>{this.props.Post}</p>}
          datetime={
            <Tooltip
              title={moment(this.props.CreatedAt).format(
                "MMMM Do YY, h:mm:ss a"
              )}
            >
              <span>
                {moment(this.props.CreatedAt).format("MMMM Do YY, h:mm:ss a")}
              </span>
            </Tooltip>
          }
        />
      </Card>
    );
  }
}

export default Post;
