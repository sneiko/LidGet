import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Request } from '../models/request';
export class RequestsService {
  constructor(private http: HttpClient) { }

  fetchAll(from: number, to: number): Observable<Request[]> {
    return this.http.get<Request[]>(`http://134.122.90.98:5000/api/requests/all?from=${from}&to=${to}`)
  }
}
