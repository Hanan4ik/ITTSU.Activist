import React from 'react';
import { View, ScrollView, Text, Image, StyleSheet, Button, } from 'react-native';

const Event = ({ images, description, onRespond, onViewComments, averageRating }) => {
  
  const displayedImages = images.slice(0, 5);

  return (
    <ScrollView style={styles.container}>
      <View style={styles.imageContainer}>
        {displayedImages.map((image, index) => (
          <Image key={index} source={{ uri: image }} style={styles.image} />
        ))}
      </View>
      <Text style={styles.description}>{description}</Text>
      <Text style={styles.rating}>Рейтинг: {averageRating.toFixed(1)} ★</Text>
      <View style={styles.buttonContainer}>
        <Button title="Откликнуться" onPress={onRespond} />
        <Button title="" onPress={onViewComments} />
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    marginTop: 15,
    padding: 10,
    width: 'auto',
    marginBottom: 15,
    backgroundColor: '#fff',
    borderRadius: 10,
    borderWidth: 2,
    borderColor: "blue",
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 1,
    },
    shadowOpacity: 0.2,
    shadowRadius: 1.41,
    elevation: 2,
  },
  imageContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-between',
  },
  image: {
    width: '100%',
    height: 1000,
    marginBottom: 5,
    borderRadius: 5,
  },
  description: {
    marginVertical: 10,
    fontSize: 28,
  },
  rating: {
    fontSize: 14,
    marginBottom: 10,
  },
  buttonContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
});

export default Event;
