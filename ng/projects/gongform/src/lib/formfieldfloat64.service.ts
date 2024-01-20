// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldFloat64, CopyFormFieldFloat64ToFormFieldFloat64DB } from './formfieldfloat64'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldFloat64Service {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldFloat64ServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldfloat64sUrl: string

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
    this.formfieldfloat64sUrl = origin + '/api/github.com/fullstack-lang/gongform/go/v1/formfieldfloat64s';
  }

  /** GET formfieldfloat64s from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB[]> {
    return this.getFormFieldFloat64s(GONG__StackPath, frontRepo)
  }
  getFormFieldFloat64s(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldFloat64DB[]>(this.formfieldfloat64sUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormFieldFloat64DB[]>('getFormFieldFloat64s', []))
      );
  }

  /** GET formfieldfloat64 by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {
    return this.getFormFieldFloat64(id, GONG__StackPath, frontRepo)
  }
  getFormFieldFloat64(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfieldfloat64sUrl}/${id}`;
    return this.http.get<FormFieldFloat64DB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldfloat64 id=${id}`)),
      catchError(this.handleError<FormFieldFloat64DB>(`getFormFieldFloat64 id=${id}`))
    );
  }

  // postFront copy formfieldfloat64 to a version with encoded pointers and post to the back
  postFront(formfieldfloat64: FormFieldFloat64, GONG__StackPath: string): Observable<FormFieldFloat64DB> {
    let formfieldfloat64DB = new FormFieldFloat64DB
    CopyFormFieldFloat64ToFormFieldFloat64DB(formfieldfloat64, formfieldfloat64DB)
    const id = typeof formfieldfloat64DB === 'number' ? formfieldfloat64DB : formfieldfloat64DB.ID
    const url = `${this.formfieldfloat64sUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldFloat64DB>(url, formfieldfloat64DB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldFloat64DB>('postFormFieldFloat64'))
    );
  }
  
  /** POST: add a new formfieldfloat64 to the server */
  post(formfieldfloat64db: FormFieldFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {
    return this.postFormFieldFloat64(formfieldfloat64db, GONG__StackPath, frontRepo)
  }
  postFormFieldFloat64(formfieldfloat64db: FormFieldFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldFloat64DB>(this.formfieldfloat64sUrl, formfieldfloat64db, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formfieldfloat64db id=${formfieldfloat64db.ID}`)
      }),
      catchError(this.handleError<FormFieldFloat64DB>('postFormFieldFloat64'))
    );
  }

  /** DELETE: delete the formfieldfloat64db from the server */
  delete(formfieldfloat64db: FormFieldFloat64DB | number, GONG__StackPath: string): Observable<FormFieldFloat64DB> {
    return this.deleteFormFieldFloat64(formfieldfloat64db, GONG__StackPath)
  }
  deleteFormFieldFloat64(formfieldfloat64db: FormFieldFloat64DB | number, GONG__StackPath: string): Observable<FormFieldFloat64DB> {
    const id = typeof formfieldfloat64db === 'number' ? formfieldfloat64db : formfieldfloat64db.ID;
    const url = `${this.formfieldfloat64sUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldFloat64DB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldfloat64db id=${id}`)),
      catchError(this.handleError<FormFieldFloat64DB>('deleteFormFieldFloat64'))
    );
  }

  // updateFront copy formfieldfloat64 to a version with encoded pointers and update to the back
  updateFront(formfieldfloat64: FormFieldFloat64, GONG__StackPath: string): Observable<FormFieldFloat64DB> {
    let formfieldfloat64DB = new FormFieldFloat64DB
    CopyFormFieldFloat64ToFormFieldFloat64DB(formfieldfloat64, formfieldfloat64DB)
    const id = typeof formfieldfloat64DB === 'number' ? formfieldfloat64DB : formfieldfloat64DB.ID
    const url = `${this.formfieldfloat64sUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormFieldFloat64DB>(url, formfieldfloat64DB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldFloat64DB>('updateFormFieldFloat64'))
    );
  }

  /** PUT: update the formfieldfloat64db on the server */
  update(formfieldfloat64db: FormFieldFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {
    return this.updateFormFieldFloat64(formfieldfloat64db, GONG__StackPath, frontRepo)
  }
  updateFormFieldFloat64(formfieldfloat64db: FormFieldFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldFloat64DB> {
    const id = typeof formfieldfloat64db === 'number' ? formfieldfloat64db : formfieldfloat64db.ID;
    const url = `${this.formfieldfloat64sUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldFloat64DB>(url, formfieldfloat64db, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formfieldfloat64db id=${formfieldfloat64db.ID}`)
      }),
      catchError(this.handleError<FormFieldFloat64DB>('updateFormFieldFloat64'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldFloat64Service', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldFloat64Service" + error); // log to console instead

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
