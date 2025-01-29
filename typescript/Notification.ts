export interface Notification {
  id: string; // UUID
  userId: string; // UUID
  isRead: boolean;
  message: string;
  createdAt: Date;
  updatedAt: Date;
}
