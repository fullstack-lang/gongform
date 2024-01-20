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

import { FormFieldSelectDB } from './formfieldselect-db'
import { FormFieldSelect, CopyFormFieldSelectToFormFieldSelectDB } from './formfieldselect'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { OptionDB } from './option-db'

@Injectable({
  providedIn: 'root'
})
export class FormFieldSelectService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldSelectServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldselectsUrl: string

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
    this.formfieldselectsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formfieldselects';
  }

  /** GET formfieldselects from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB[]> {
    return this.getFormFieldSelects(GONG__StackPath, frontRepo)
  }
  getFormFieldSelects(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldSelectDB[]>(this.formfieldselectsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormFieldSelectDB[]>('getFormFieldSelects', []))
      );
  }

  /** GET formfieldselect by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {
    return this.getFormFieldSelect(id, GONG__StackPath, frontRepo)
  }
  getFormFieldSelect(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfieldselectsUrl}/${id}`;
    return this.http.get<FormFieldSelectDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldselect id=${id}`)),
      catchError(this.handleError<FormFieldSelectDB>(`getFormFieldSelect id=${id}`))
    );
  }

  /** POST: add a new formfieldselect to the server */
  post(formfieldselectdb: FormFieldSelectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {
    return this.postFormFieldSelect(formfieldselectdb, GONG__StackPath, frontRepo)
  }
  postFormFieldSelect(formfieldselectdb: FormFieldSelectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    if (formfieldselectdb.Value != undefined) {
      formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Int64 = formfieldselectdb.Value.ID
      formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Valid = true
    }
    formfieldselectdb.Value = undefined
    formfieldselectdb.FormFieldSelectPointersEncoding.Options = []
    for (let _option of formfieldselectdb.Options) {
      formfieldselectdb.FormFieldSelectPointersEncoding.Options.push(_option.ID)
    }
    formfieldselectdb.Options = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldSelectDB>(this.formfieldselectsUrl, formfieldselectdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formfieldselectdb.Value = frontRepo.Options.get(formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Int64)
        formfieldselectdb.Options = new Array<OptionDB>()
        for (let _id of formfieldselectdb.FormFieldSelectPointersEncoding.Options) {
          let _option = frontRepo.Options.get(_id)
          if (_option != undefined) {
            formfieldselectdb.Options.push(_option!)
          }
        }
        // this.log(`posted formfieldselectdb id=${formfieldselectdb.ID}`)
      }),
      catchError(this.handleError<FormFieldSelectDB>('postFormFieldSelect'))
    );
  }

  /** DELETE: delete the formfieldselectdb from the server */
  delete(formfieldselectdb: FormFieldSelectDB | number, GONG__StackPath: string): Observable<FormFieldSelectDB> {
    return this.deleteFormFieldSelect(formfieldselectdb, GONG__StackPath)
  }
  deleteFormFieldSelect(formfieldselectdb: FormFieldSelectDB | number, GONG__StackPath: string): Observable<FormFieldSelectDB> {
    const id = typeof formfieldselectdb === 'number' ? formfieldselectdb : formfieldselectdb.ID;
    const url = `${this.formfieldselectsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldSelectDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldselectdb id=${id}`)),
      catchError(this.handleError<FormFieldSelectDB>('deleteFormFieldSelect'))
    );
  }

  // updateFront copy formfieldselect to a version with encoded pointers and update to the back
  updateFront(formfieldselect: FormFieldSelect, GONG__StackPath: string): Observable<FormFieldSelectDB> {
    let formfieldselectDB = new FormFieldSelectDB
    CopyFormFieldSelectToFormFieldSelectDB(formfieldselect, formfieldselectDB)
    const id = typeof formfieldselectDB === 'number' ? formfieldselectDB : formfieldselectDB.ID
    const url = `${this.formfieldselectsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormFieldSelectDB>(url, formfieldselectDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldSelectDB>('updateFormFieldSelect'))
    );
  }

  /** PUT: update the formfieldselectdb on the server */
  update(formfieldselectdb: FormFieldSelectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {
    return this.updateFormFieldSelect(formfieldselectdb, GONG__StackPath, frontRepo)
  }
  updateFormFieldSelect(formfieldselectdb: FormFieldSelectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldSelectDB> {
    const id = typeof formfieldselectdb === 'number' ? formfieldselectdb : formfieldselectdb.ID;
    const url = `${this.formfieldselectsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    if (formfieldselectdb.Value != undefined) {
      formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Int64 = formfieldselectdb.Value.ID
      formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Valid = true
    }
    formfieldselectdb.Value = undefined
    formfieldselectdb.FormFieldSelectPointersEncoding.Options = []
    for (let _option of formfieldselectdb.Options) {
      formfieldselectdb.FormFieldSelectPointersEncoding.Options.push(_option.ID)
    }
    formfieldselectdb.Options = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldSelectDB>(url, formfieldselectdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formfieldselectdb.Value = frontRepo.Options.get(formfieldselectdb.FormFieldSelectPointersEncoding.ValueID.Int64)
        formfieldselectdb.Options = new Array<OptionDB>()
        for (let _id of formfieldselectdb.FormFieldSelectPointersEncoding.Options) {
          let _option = frontRepo.Options.get(_id)
          if (_option != undefined) {
            formfieldselectdb.Options.push(_option!)
          }
        }
        // this.log(`updated formfieldselectdb id=${formfieldselectdb.ID}`)
      }),
      catchError(this.handleError<FormFieldSelectDB>('updateFormFieldSelect'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldSelectService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldSelectService" + error); // log to console instead

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
