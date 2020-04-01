import React, { Component } from "react";
import { Affix, Button, Modal, Input } from "antd";
import { connect } from "react-redux";
import { PlusOutlined } from "@ant-design/icons";
import { addPostProcess } from "../../actions/actions";

const { TextArea } = Input;

class AddPost extends Component {
  constructor(props) {
    super(props);
    this.state = {
      visible: false,
      confirmLoading: false,
      postText: ""
    };
  }
  showModal = () => {
    this.setState({
      visible: true
    });
  };
  handleOk = () => {
    this.setState({
      ModalText: "Posting...",
      confirmLoading: true
    });
    this.props.addPostProcess(this.props.token, this.props.username, this.state.postText);
    setTimeout(() => {
      this.setState({
        visible: false,
        confirmLoading: false
      });
    }, 2000);
  };

  handleCancel = () => {
    console.log("Clicked cancel button");
    this.setState({
      visible: false
    });
  };
  render() {
    return (
      <div>
        <Modal
          visible={this.state.visible}
          title="Add a post"
          confirmLoading={this.state.confirmLoading}
          onOk={this.handleOk}
          onCancel={this.handleCancel}
        >
            <TextArea rows={4} onChange={event => this.setState({postText: event.target.value})} />
        </Modal>
        <Affix offsetBottom={150}>
          <Button
            type="primary"
            shape="circle"
            icon={<PlusOutlined />}
            style={{ float: "right", marginRight: "10%" }}
            size="large"
            onClick={this.showModal}
          />
        </Affix>
      </div>
    );
  }
}

const mapStateToProps = state => ({
    username: state.me.username,
    token: state.token
  });

  export default connect(mapStateToProps, { addPostProcess })(AddPost);
