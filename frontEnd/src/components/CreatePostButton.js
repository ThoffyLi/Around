import React from "react";
import { Button, Modal, message } from "antd";
import { CreatePostForm } from "./CreatePostForm";
import {
  API_ROOT,
  AUTH_HEADER,
  POS_KEY,
  TOKEN_KEY,
  LOC_SHAKE
} from "../constants";

export class CreatePostButton extends React.Component {
  state = {
    visible: false,
    confirmLoading: false
  };

  getFormRef = formInstance => {
    this.form = formInstance;
  };

  showModal = () => {
    this.setState({
      visible: true
    });
  };

  handleOk = () => {
    this.form.validateFields((err, values) => {
      if (!err) {
        const token = localStorage.getItem(TOKEN_KEY);
        const { lat, lon } = JSON.parse(localStorage.getItem(POS_KEY));
        const formData = new FormData();

        formData.append("message", values.message);
        formData.append("image", values.imageorvideo[0].originFileObj);
        formData.append("lat", lat + 2 * Math.random() * LOC_SHAKE - LOC_SHAKE);
        formData.append("lon", lon + 2 * Math.random() * LOC_SHAKE - LOC_SHAKE);
        formData.append("lon", lon + 2 * Math.random() * LOC_SHAKE - LOC_SHAKE);
        //console.log(formData.get("image"));
        this.setState({
          confirmLoading: true
        });

        //fire api call
        //json sets binary image file has problem
        fetch(`${API_ROOT}/post`, {
          headers: {
            Authorization: `${AUTH_HEADER} ${token}`
          },
          method: "POST",
          body: formData
        })
          .then(response => {
            if (response.ok) {
              this.form.resetFields();
              this.setState({
                confirmLoading: false,
                visible: false
              });
              this.props.loadNearbyPosts();
              return response;
            }
            throw new Error(response.statusText);
          })
          .then(() => {
            message.success("Post created successfully!");
          })
          .catch(err => {
            // console.log(err);
            message.error("Failed to create the post.");
            this.setState({
              confirmLoading: false
            });
          });
      }
    });
  };

  handleCancel = () => {
    console.log("Clicked cancel button");
    this.setState({
      visible: false
    });
  };

  render() {
    const { visible, confirmLoading } = this.state;
    return (
      <div>
        <Button type="primary" onClick={this.showModal}>
          Create New Post
        </Button>
        <Modal
          title="Create New Post"
          visible={visible}
          onOk={this.handleOk}
          confirmLoading={confirmLoading}
          onCancel={this.handleCancel}
          okText="Create"
        >
          <CreatePostForm ref={this.getFormRef} />
        </Modal>
      </div>
    );
  }
}
