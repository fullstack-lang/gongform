// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { FormFieldTimeDB } from './formfieldtime-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldTimeService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldTimeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldtimesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.formfieldtimesUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formfieldtimes';
  }

  /** GET formfieldtimes from the server */
  getFormFieldTimes(GONG__StackPath: string): Observable<FormFieldTimeDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldTimeDB[]>(this.formfieldtimesUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formfieldtimes')),
        catchError(this.handleError<FormFieldTimeDB[]>('getFormFieldTimes', []))
      );
  }

  /** GET formfieldtime by id. Will 404 if id not found */
  getFormFieldTime(id: number, GONG__StackPath: string): Observable<FormFieldTimeDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfieldtimesUrl}/${id}`;
    return this.http.get<FormFieldTimeDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldtime id=${id}`)),
      catchError(this.handleError<FormFieldTimeDB>(`getFormFieldTime id=${id}`))
    );
  }

  /** POST: add a new formfieldtime to the server */
  postFormFieldTime(formfieldtimedb: FormFieldTimeDB, GONG__StackPath: string): Observable<FormFieldTimeDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldTimeDB>(this.formfieldtimesUrl, formfieldtimedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formfieldtimedb id=${formfieldtimedb.ID}`)
      }),
      catchError(this.handleError<FormFieldTimeDB>('postFormFieldTime'))
    );
  }

  /** DELETE: delete the formfieldtimedb from the server */
  deleteFormFieldTime(formfieldtimedb: FormFieldTimeDB | number, GONG__StackPath: string): Observable<FormFieldTimeDB> {
    const id = typeof formfieldtimedb === 'number' ? formfieldtimedb : formfieldtimedb.ID;
    const url = `${this.formfieldtimesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldTimeDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldtimedb id=${id}`)),
      catchError(this.handleError<FormFieldTimeDB>('deleteFormFieldTime'))
    );
  }

  /** PUT: update the formfieldtimedb on the server */
  updateFormFieldTime(formfieldtimedb: FormFieldTimeDB, GONG__StackPath: string): Observable<FormFieldTimeDB> {
    const id = typeof formfieldtimedb === 'number' ? formfieldtimedb : formfieldtimedb.ID;
    const url = `${this.formfieldtimesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldTimeDB>(url, formfieldtimedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formfieldtimedb id=${formfieldtimedb.ID}`)
      }),
      catchError(this.handleError<FormFieldTimeDB>('updateFormFieldTime'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldTimeService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldTimeService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
      console.log(message)
  }
}
