import React, { useState, useEffect } from 'react';
import { View, Text, Image, StyleSheet, Dimensions } from 'react-native';
import BeautifulButton from './beautifulButton';


const Event = ({ images, description, onRespond, onViewComments, averageRating }) => {
  const [aspectRatio, setAspectRatio] = useState(16 / 9);
  const displayImage = images[0];

  useEffect(() => {
    if (displayImage) {
      // Fetch the actual image dimensions from the URL
      Image.getSize(displayImage, (width, height) => {
        setAspectRatio(width / height);
      }, (error) => {
        console.warn("Couldn't get image size:", error);
      });
    }
  }, [displayImage]);

  return (
    <View style={styles.container}>
      <View style={styles.imageWrapper}>
        {displayImage && (
          <Image 
            source={{ uri: displayImage }} 
            style={[
              styles.image, 
              { aspectRatio: aspectRatio }
            ]}
          />
        )}
      </View>

      <View style={styles.content}>
        <Text style={styles.rating}>Рейтинг: {averageRating.toFixed(1)} ★</Text>
        <Text style={styles.description}>{description}</Text>
        
        <View style={styles.buttonContainer}>
          <BeautifulButton
            btnProps={{ color: '#007AFF', onPress: onRespond }}
            btnTextProps={{ color: "#fff", fontSize: 18, text: 'Откликнуться' }}
          />
          <BeautifulButton
            btnProps={{ color: '#007AFF', onPress: onViewComments }}
            btnTextProps={{ color: "#fff", fontSize: 18, text: 'Открыть комментарии' }}
          />
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    marginTop: 15,
    marginHorizontal: '5%',
    backgroundColor: '#fff',
    borderRadius: 12,
    borderWidth: 1,
    borderColor: "#E0E0E0",
    overflow: 'hidden',
    elevation: 3,
  },
  imageWrapper: {
    width: '100%',
    backgroundColor: '#f0f0f0',
  },
  image: {
    width: '100%',
    resizeMode: 'contain', 
  },
  content: {
    padding: 15,
  },
  description: {
    marginBottom: 15,
    fontSize: 18,
    color: '#333',
  },
  rating: {
    fontSize: 14,
    color: '#666',
    fontWeight: 'bold',
    marginBottom: 5,
  },
  buttonContainer: {
    marginTop: 5,
  }
});

export default Event;