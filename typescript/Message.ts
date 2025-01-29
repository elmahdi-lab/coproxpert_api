
export interface Message {
  id: string; // UUID
  content: string;
  fromId: string; // UUID
  fromUser?: User;
  toId: string; // UUID
  toUser?: User;
  isRead: boolean;
  createdAt: Date;
  updatedAt: Date;
}
