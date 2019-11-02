import React from "react";
import GridGallery from "react-grid-gallery";
import PropTypes from "prop-types";

export class Gallery extends React.Component {
  static propTypes = {
    images: PropTypes.arrayOf(
      PropTypes.shape({
        user: PropTypes.string.isRequired,
        src: PropTypes.string.isRequired,
        thumbnail: PropTypes.string.isRequired,
        caption: PropTypes.string,
        thumbnailWidth: PropTypes.number.isRequired,
        thumbnailHeight: PropTypes.number.isRequired
      })
    ).isRequired
  };

  render() {
    var images = this.props.images.map(i => {
      const customOverlay = (
        <div className="caption">
          <div>{`${i.user}: ${i.caption}`}</div>
        </div>
      );
      return { ...i, customOverlay };
    });

    return (
      <div className="wrapper">
        <GridGallery
          backdropClosesModal
          images={images}
          enableImageSelection={false}
        />
      </div>
    );
  }
}
