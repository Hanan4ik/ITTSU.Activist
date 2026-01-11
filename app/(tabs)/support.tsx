import React, { useState } from 'react';
import {
  View,
  Text,
  TextInput,
  StyleSheet,
  TouchableOpacity,
  Alert,
  ActivityIndicator,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TouchableWithoutFeedback,
  Keyboard,
  Pressable
} from 'react-native';

const API_URL = 'http://127.0.0.1:8000';

export default function QuestionScreen() {
  // --- State ---
  const [email, setEmail] = useState('');
  const [question, setQuestion] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  // --- Logic ---
  const handleSubmit = async () => {
    // 1. Basic Validation
    if (!email.trim() || !question.trim()) {
      Alert.alert('Missing Fields', 'Please fill in both your email and your question.');
      return;
    }

    // Simple Email Regex Check
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      Alert.alert('Invalid Email', 'Please enter a valid email address.');
      return;
    }

    setIsSubmitting(true);

    try {
      // 2. Send Data to Backend
      const response = await fetch(`${API_URL}/api/askQuestion`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          message: question,
        }),
      });

      if (response.ok) {
        Alert.alert('Success', 'Your question has been sent! We will contact you soon.');
        // Clear form
        setEmail('');
        setQuestion('');
      } else {
        const errorData = await response.json();
        Alert.alert('Error', errorData.message || 'Something went wrong.');
      }
    } catch (error) {
      console.error(error);
      Alert.alert('Network Error', 'Could not reach the server.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={styles.container}
      >
          <ScrollView contentContainerStyle={styles.scrollContainer}>
            
            <View style={styles.headerContainer}>
              <Text style={styles.title}>Что-то не понятно?</Text>
              <Text style={styles.subtitle}>
                Оставьте контактные данные и напишите ваш вопрос. Мы непременно вам ответим.
              </Text>
            </View>

            <View style={styles.formCard}>
              {/* Email Input */}
              <Text style={styles.label}>Email</Text>
              <TextInput
                style={styles.input}
                placeholder="user@example.com"
                value={email}
                onChangeText={setEmail}
                keyboardType="email-address"
              />

              
              <Text style={styles.label}>Ваш вопрос</Text>
              <TextInput
                style={[styles.input, styles.textArea]}
                placeholder="Пишите вопрос здесь..."
                value={question}
                onChangeText={setQuestion}
                multiline={true}
                numberOfLines={6}
                textAlignVertical="top"
              />

              {/* Submit Button */}
              <TouchableOpacity
                style={styles.submitButton}
                onPress={handleSubmit}
                disabled={isSubmitting}
              >
                {isSubmitting ? (
                  <ActivityIndicator color="#fff" />
                ) : (
                  <Text style={styles.submitButtonText}>Отправить</Text>
                )}
              </TouchableOpacity>
            </View>

          </ScrollView>
      </KeyboardAvoidingView>
    
    
  );
}

// --- Styles ---
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  scrollContainer: {
    padding: 20,
    flexGrow: 1,
  },
  headerContainer: {
    marginBottom: 25,
    alignItems: 'center',
  },
  title: {
    fontSize: 26,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 10,
  },
  subtitle: {
    fontSize: 16,
    color: '#666',
    textAlign: 'center',
    lineHeight: 22,
  },
  formCard: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  label: {
    fontSize: 16,
    fontWeight: '600',
    color: '#333',
    marginBottom: 8,
    marginTop: 5,
  },
  input: {
    backgroundColor: '#fafafa',
    borderWidth: 1,
    borderColor: '#e0e0e0',
    borderRadius: 8,
    padding: 12,
    fontSize: 16,
    color: '#333',
    marginBottom: 20,
  },
  textArea: {
    height: 150,
    paddingTop: 12,
  },
  submitButton: {
    backgroundColor: '#007AFF',
    paddingVertical: 15,
    borderRadius: 8,
    alignItems: 'center',
    marginTop: 10,
  },
  submitButtonText: {
    color: '#fff',
    fontSize: 18,
    fontWeight: 'bold',
  },
});