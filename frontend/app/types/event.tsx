import React from 'react';
import { View, ScrollView, Text, Image, StyleSheet, Button, Dimensions } from 'react-native';

const Event = ({ images, description, onRespond, onViewComments, averageRating }) => {
  images = images.slice(0, 1); // В будущем придумать, как группировать изображения красиво как в телеге
  return (
    <ScrollView style={styles.container}>
      <View style={styles.imageContainer}>
        {
        images.map((image, index) => ( 
          
          <Image key={index} source={{ uri: image }} style={styles.image}/>
        ))}
      </View>
      <Text style={styles.rating}>Рейтинг: {averageRating.toFixed(1)} ★</Text>
      <Text style={styles.description}>{description}</Text>
      
      <View style={styles.buttonContainer}>
        <Button title="Откликнуться" onPress={onRespond} style={styles.rating} />
        <View style={styles.spacer} />
        <Button title="Открыть комментарии" onPress={onViewComments} />
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    marginTop: 15,
    marginLeft: '5%',
    marginRight: '5%',
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
    
  },
   image: {
    width: '100%',
    height: (Dimensions.get('screen').height),
    
    resizeMode: 'contain'
  },
  imageContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  description: {
    marginVertical: 10,
    fontSize: 29,
  },
  rating: {
    fontSize: 14,
    marginBottom: 10,
  },
  buttonContainer: {
    flexDirection: 'column',
    justifyContent: 'flex-start',
  },
  spacer: {
    marginVertical: 10,
  }
});

export default Event;
