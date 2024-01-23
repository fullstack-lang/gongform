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

import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldInt, CopyFormFieldIntToFormFieldIntDB } from './formfieldint'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldIntService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldIntServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldintsUrl: string

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
    this.formfieldintsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formfieldints';
  }

  /** GET formfieldints from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB[]> {
    return this.getFormFieldInts(GONG__StackPath, frontRepo)
  }
  getFormFieldInts(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldIntDB[]>(this.formfieldintsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormFieldIntDB[]>('getFormFieldInts', []))
      );
  }

  /** GET formfieldint by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {
    return this.getFormFieldInt(id, GONG__StackPath, frontRepo)
  }
  getFormFieldInt(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfieldintsUrl}/${id}`;
    return this.http.get<FormFieldIntDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldint id=${id}`)),
      catchError(this.handleError<FormFieldIntDB>(`getFormFieldInt id=${id}`))
    );
  }

  // postFront copy formfieldint to a version with encoded pointers and post to the back
  postFront(formfieldint: FormFieldInt, GONG__StackPath: string): Observable<FormFieldIntDB> {
    let formfieldintDB = new FormFieldIntDB
    CopyFormFieldIntToFormFieldIntDB(formfieldint, formfieldintDB)
    const id = typeof formfieldintDB === 'number' ? formfieldintDB : formfieldintDB.ID
    const url = `${this.formfieldintsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldIntDB>(url, formfieldintDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldIntDB>('postFormFieldInt'))
    );
  }
  
  /** POST: add a new formfieldint to the server */
  post(formfieldintdb: FormFieldIntDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {
    return this.postFormFieldInt(formfieldintdb, GONG__StackPath, frontRepo)
  }
  postFormFieldInt(formfieldintdb: FormFieldIntDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldIntDB>(this.formfieldintsUrl, formfieldintdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formfieldintdb id=${formfieldintdb.ID}`)
      }),
      catchError(this.handleError<FormFieldIntDB>('postFormFieldInt'))
    );
  }

  /** DELETE: delete the formfieldintdb from the server */
  delete(formfieldintdb: FormFieldIntDB | number, GONG__StackPath: string): Observable<FormFieldIntDB> {
    return this.deleteFormFieldInt(formfieldintdb, GONG__StackPath)
  }
  deleteFormFieldInt(formfieldintdb: FormFieldIntDB | number, GONG__StackPath: string): Observable<FormFieldIntDB> {
    const id = typeof formfieldintdb === 'number' ? formfieldintdb : formfieldintdb.ID;
    const url = `${this.formfieldintsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldIntDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldintdb id=${id}`)),
      catchError(this.handleError<FormFieldIntDB>('deleteFormFieldInt'))
    );
  }

  // updateFront copy formfieldint to a version with encoded pointers and update to the back
  updateFront(formfieldint: FormFieldInt, GONG__StackPath: string): Observable<FormFieldIntDB> {
    let formfieldintDB = new FormFieldIntDB
    CopyFormFieldIntToFormFieldIntDB(formfieldint, formfieldintDB)
    const id = typeof formfieldintDB === 'number' ? formfieldintDB : formfieldintDB.ID
    const url = `${this.formfieldintsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormFieldIntDB>(url, formfieldintDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldIntDB>('updateFormFieldInt'))
    );
  }

  /** PUT: update the formfieldintdb on the server */
  update(formfieldintdb: FormFieldIntDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {
    return this.updateFormFieldInt(formfieldintdb, GONG__StackPath, frontRepo)
  }
  updateFormFieldInt(formfieldintdb: FormFieldIntDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldIntDB> {
    const id = typeof formfieldintdb === 'number' ? formfieldintdb : formfieldintdb.ID;
    const url = `${this.formfieldintsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldIntDB>(url, formfieldintdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formfieldintdb id=${formfieldintdb.ID}`)
      }),
      catchError(this.handleError<FormFieldIntDB>('updateFormFieldInt'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldIntService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldIntService" + error); // log to console instead

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
