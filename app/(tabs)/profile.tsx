import React, { useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  Modal,
  TextInput,
  Alert,
  ActivityIndicator,
  SafeAreaView,
  ScrollView,
} from 'react-native';


class User {
  id: number;
  tableNumber: number;
  username: string;
  email: string;
  rights: number;

  constructor(id: number, tableNumber: number, username:string, email:string, rights:number){
    this.id = id;
    this.tableNumber = tableNumber;
    this.username = username;
    this.email = email;
    this.rights = rights;
  }
}

const API_URL = 'http://127.0.0.1:8000';

export default function ProfileScreen() {
  
  const [user, setUser] = useState<User | null>(null);
  const [modalVisible, setModalVisible] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  // Form State
  const [loginInput, setLoginInput] = useState('');
  const [passwordInput, setPasswordInput] = useState('');

  // --- Logic ---

  const handleLogin = async () => {
    if (!loginInput || !passwordInput) {
      Alert.alert('Error', 'Please fill in both fields');
      return;
    }

    setIsLoading(true);

    try {
      const response = await fetch(`${API_URL}/api/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          login: loginInput,
          password: passwordInput,
        }),
      });
      console.log(response);

      if (response.ok) {
        

        setModalVisible(false);
        setLoginInput('');
        setPasswordInput('');
      } else {
        Alert.alert('Login Failed', data.message || 'Invalid credentials');
      }
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'Network request failed');
    } finally {
      setIsLoading(false);
    }
  };

  const handleLogout = () => {
    setUser(null);
  };

  // --- Render Helpers ---

  const renderUserInfoRow = (label: string, value: string | number) => (
    <View style={styles.infoRow}>
      <Text style={styles.infoLabel}>{label}:</Text>
      <Text style={styles.infoValue}>{value}</Text>
    </View>
  );

  return (
    <View style={styles.container}>
      <ScrollView contentContainerStyle={styles.contentContainer}>

        {user ? (
          // --- View: Logged In ---
          <View style={styles.card}>
            <View style={styles.avatarPlaceholder}>
              <Text style={styles.avatarText}>{user.username.charAt(0).toUpperCase()}</Text>
            </View>
            
            <Text style={styles.welcomeText}>Welcome, {user.username}!</Text>

            <View style={styles.infoContainer}>
              {renderUserInfoRow('ID', user.id)}
              {renderUserInfoRow('Table Number', user.tableNumber)}
              {renderUserInfoRow('Email', user.email)}
              {renderUserInfoRow('Access Rights', user.rights)}
            </View>

            <TouchableOpacity style={styles.logoutButton} onPress={handleLogout}>
              <Text style={styles.logoutButtonText}>Log Out</Text>
            </TouchableOpacity>
          </View>
        ) : (
          // --- View: Logged Out ---
          <View style={styles.guestContainer}>
            <Text style={styles.guestText}>Вы не авторизованы.</Text>
            <Text style={styles.subGuestText}>Войдите, чтобы посмотреть свой профиль.</Text>
            
            <TouchableOpacity 
              style={styles.loginButton} 
              onPress={() => setModalVisible(true)}
            >
              <Text style={styles.loginButtonText}>Open Login</Text>
            </TouchableOpacity>
          </View>
        )}
      </ScrollView>

      {/* --- Login Modal --- */}
      <Modal
        animationType='fade'
        transparent={true}
        visible={modalVisible}
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalOverlay}>
          <View style={styles.modalView}>
            <Text style={styles.modalTitle}>Login</Text>

            <TextInput
              style={styles.input}
              placeholder="Username"
              value={loginInput}
              onChangeText={setLoginInput}
              autoCapitalize="none"
            />
            
            <TextInput
              style={styles.input}
              placeholder="Password"
              value={passwordInput}
              onChangeText={setPasswordInput}
              secureTextEntry
            />

            <View style={styles.modalButtons}>
              <TouchableOpacity 
                style={[styles.modalBtn, styles.cancelBtn]} 
                onPress={() => setModalVisible(false)}
              >
                <Text style={styles.btnTextBlack}>Cancel</Text>
              </TouchableOpacity>

              <TouchableOpacity 
                style={[styles.modalBtn, styles.submitBtn]} 
                onPress={handleLogin}
                disabled={isLoading}
              >
                {isLoading ? (
                  <ActivityIndicator color="#fff" />
                ) : (
                  <Text style={styles.btnTextWhite}>Login</Text>
                )}
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </Modal>
    </View>
  );
}

// --- Styles ---
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  contentContainer: {
    padding: 20,
    alignItems: 'center',
  },
  headerTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
    color: '#333',
  },
  // Guest Styles
  guestContainer: {
    alignItems: 'center',
    marginTop: 50,
  },
  guestText: {
    fontSize: 18,
    color: '#666',
  },
  subGuestText: {
    fontSize: 14,
    color: '#999',
    marginBottom: 20,
  },
  loginButton: {
    backgroundColor: '#007AFF',
    paddingHorizontal: 30,
    paddingVertical: 12,
    borderRadius: 8,
  },
  loginButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  // User Profile Styles
  card: {
    backgroundColor: '#fff',
    width: '100%',
    borderRadius: 12,
    padding: 20,
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  avatarPlaceholder: {
    width: 80,
    height: 80,
    borderRadius: 40,
    backgroundColor: '#e1e1e1',
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 15,
  },
  avatarText: {
    fontSize: 32,
    fontWeight: 'bold',
    color: '#555',
  },
  welcomeText: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  infoContainer: {
    width: '100%',
    marginBottom: 20,
  },
  infoRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#f0f0f0',
  },
  infoLabel: {
    fontWeight: '600',
    color: '#555',
  },
  infoValue: {
    color: '#333',
  },
  logoutButton: {
    marginTop: 10,
    paddingVertical: 10,
  },
  logoutButtonText: {
    color: '#FF3B30',
    fontSize: 16,
  },
  // Modal Styles
  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0,0,0,0.5)',
    justifyContent: 'center',
    alignItems: 'center',
  },
  modalView: {
    width: '85%',
    backgroundColor: 'white',
    borderRadius: 20,
    padding: 25,
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  modalTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 15,
  },
  input: {
    width: '100%',
    height: 50,
    borderWidth: 1,
    borderColor: '#ddd',
    borderRadius: 8,
    paddingHorizontal: 15,
    marginBottom: 15,
    backgroundColor: '#fafafa',
  },
  modalButtons: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
    marginTop: 10,
  },
  modalBtn: {
    flex: 1,
    padding: 12,
    borderRadius: 8,
    alignItems: 'center',
    marginHorizontal: 5,
  },
  cancelBtn: {
    backgroundColor: '#e0e0e0',
  },
  submitBtn: {
    backgroundColor: '#007AFF',
  },
  btnTextWhite: {
    color: '#fff',
    fontWeight: 'bold',
  },
  btnTextBlack: {
    color: '#333',
    fontWeight: 'bold',
  },
});