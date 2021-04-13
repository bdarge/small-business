import { Customer } from './customer';
import { User } from './user';
import { QuoteDetail } from './quoteDetail';
import { QuoteItem } from './quoteItem';
import { QuoteSummary} from './quoteSummary';
import {Transaction} from './transaction';

export interface Quote extends Transaction{
  quoteNumber?: number
  quoteItems?: QuoteItem[]
  quoteSummary: string
}
