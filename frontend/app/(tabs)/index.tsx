import React from 'react';
import { 
  StyleSheet, 
  ScrollView, 
  View, 
   
  StatusBar 
} from 'react-native';
import EventComponent from "../types/event"; // Renamed to avoid conflict with standard 'Event' naming

const API_URL = "http://127.0.0.1:8000";

// Mock Data - In a real app, this might come from a fetch() call
const events = [
  {
    id: 0,
    images: [
      "https://eventum-premo.ru/sites/default/files/wiki/925e9c4cf9b7a9ee95214fb1df955efce2c80384.jpg",
      "https://antonenko-media.com/wp-content/uploads/2023/04/photo-event01845.jpg",
    ],
    description: "Нижний пост хуйня, приходите к нам у нас напитки, девочки и блэкджек. АОЫРАЛОРЫАОШРЫШОВАРФШОВРФЫШОРВФОШЛРВЛОФРВЛОФЫРВЛОФРВЛОФРВЛОФЫРВЛОФЫРВЛОФЫРВЛОФЫРВЛОФРВЛоФР",
    averageRating: 4.8,
  },
  {
    id: 1,
    images: [
      "https://antonenko-media.com/wp-content/uploads/2023/04/photo-event01845.jpg",
      "https://eventum-premo.ru/sites/default/files/wiki/925e9c4cf9b7a9ee95214fb1df955efce2c80384.jpg",
    ],
    description: "Тут можно много смотреть, приходите, пожалуйста и ставьте классы ЩЛАОЛОЫВРАЫЛОВРАЛОЫВРАЛОЫВРАЛОЫРВАОЛИЫРОВРАГИЕНГМАМАМЕ?НАЫФПВОФЫИВЛОФЫВЛОФЛОВРФЛОВРФОЛРВОЛР",
    averageRating: 4.2,
  }
];

export default function EventsPage() {
  
  const handleRespond = (id: string) => {
    console.log(`Responded to event: ${id}`);
  };

  const handleViewComments = (id: string) => {
    console.log(`Viewing comments for event: ${id}`);
  };

  return (
    <View style={styles.container}>
      <StatusBar barStyle="dark-content" />

      {/* Single ScrollView for better performance */}
      <ScrollView 
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.scrollContent}
      >
        {events.map((item) => (
          <View key={item.id} style={styles.eventWrapper}>
            <EventComponent
              images={item.images}
              description={item.description}
              averageRating={item.averageRating}
              onRespond={() => handleRespond(item.id)}
              onViewComments={() => handleViewComments(item.id)}
            />
          </View>
        ))}
        
        <View style={{ height: 40 }} />
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f8f9fa', // Light grey background
  },
  header: {
    paddingHorizontal: 20,
    paddingVertical: 15,
    backgroundColor: '#fff',
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },
  headerTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#1a1a1a',
  },
  scrollContent: {
    paddingVertical: 10,
  },
  eventWrapper: {
    marginBottom: 20, // Space between event cards
    // If your Event component doesn't have internal padding, add horizontal padding here
    paddingHorizontal: 15, 
  },
  
});