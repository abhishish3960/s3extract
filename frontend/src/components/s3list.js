import React, { useEffect, useState } from 'react';
import axios from 'axios';

const S3ContentDisplay = () => {
  const [images, setImages] = useState([]);
  const [texts, setTexts] = useState([]);

  useEffect(() => {
    const fetchContent = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/contents');
        setImages(response.data.images);
        setTexts(response.data.texts);
      } catch (error) {
        console.error('Error fetching content', error);
      }
    };

    fetchContent();
  }, []);

  return (
    <div>
      <h1>Images</h1>
      {images.length > 0 ? (
        images.map((item, index) => (
          <div key={index}>
            <h2>{item.name}</h2>
            <img src={item.url} alt={`Image ${index}`} style={{ width: '200px', height: 'auto' }} />
          </div>
        ))
      ) : (
        <p>No images found.</p>
      )}
      <h1>Texts</h1>
      {texts.length > 0 ? (
        texts.map((item, index) => (
          <div key={index}>
            <h2>{item.name}</h2>
            <p>{item.content}</p>
            <a href={item.url} target="_blank" rel="noopener noreferrer">Download Text File</a>
          </div>
        ))
      ) : (
        <p>No texts found.</p>
      )}
    </div>
  );
};

export default S3ContentDisplay;
