export interface User {
  id: string; // UUID
  username?: string;
  firstName?: string;
  lastName?: string;
  phoneNumber?: string;
  address?: string;
  city?: string;
  province?: string;
  zipCode?: string;
  country?: string;
  isClaimed?: boolean;
  isEmailVerified?: boolean;
  isPhoneVerified?: boolean;
  tries?: number;
  lockExpiresAt?: Date;
  passwordResetToken?: string; // UUID
  resetTokenExpiresAt?: Date;
  password?: string;
  refreshToken: string; // UUID
  refreshTokenExpiresAt?: Date;
  token?: string;
  signInProvider?: string;
  providerId?: string;
  createdAt: Date;
  updatedAt: Date;
}
