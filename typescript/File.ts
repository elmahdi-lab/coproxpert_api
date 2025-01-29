export type Provider = 'aws' | 'gcp';

export interface File {
  id: string; // UUID
  uploaderId: string; // UUID
  bucketName: string;
  provider: Provider;
  fileType: string;
  publicUrl: string;
  privateUrl: string;
  createdAt: Date;
  updatedAt: Date;
}
