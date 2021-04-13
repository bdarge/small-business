import {Customer} from './customer';
import {OrderItem} from './orderItem';
import {User} from './user';
import {Transaction} from './transaction';

export interface Order extends Transaction{
  comment: string;
  deliveryDate: string;
  invoiceNumber?: number;
  purchaseOrderNumber?: number;
  orderItems?: OrderItem[];
  currency: string;
}
