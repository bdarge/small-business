import {User} from './user';
import {Customer} from './customer';

export interface Transaction {
  id: string;
  description: string;
  createdAt: string;
  type: string;
  user: User;
  customer: Customer;
}
