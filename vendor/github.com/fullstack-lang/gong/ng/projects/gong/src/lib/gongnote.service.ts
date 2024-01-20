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

import { GongNoteDB } from './gongnote-db'
import { GongNote, CopyGongNoteToGongNoteDB } from './gongnote'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { GongLinkDB } from './gonglink-db'

@Injectable({
  providedIn: 'root'
})
export class GongNoteService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongNoteServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongnotesUrl: string

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
    this.gongnotesUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gongnotes';
  }

  /** GET gongnotes from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB[]> {
    return this.getGongNotes(GONG__StackPath, frontRepo)
  }
  getGongNotes(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<GongNoteDB[]>(this.gongnotesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<GongNoteDB[]>('getGongNotes', []))
      );
  }

  /** GET gongnote by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {
    return this.getGongNote(id, GONG__StackPath, frontRepo)
  }
  getGongNote(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.gongnotesUrl}/${id}`;
    return this.http.get<GongNoteDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gongnote id=${id}`)),
      catchError(this.handleError<GongNoteDB>(`getGongNote id=${id}`))
    );
  }

  // postFront copy gongnote to a version with encoded pointers and post to the back
  postFront(gongnote: GongNote, GONG__StackPath: string): Observable<GongNoteDB> {
    let gongnoteDB = new GongNoteDB
    CopyGongNoteToGongNoteDB(gongnote, gongnoteDB)
    const id = typeof gongnoteDB === 'number' ? gongnoteDB : gongnoteDB.ID
    const url = `${this.gongnotesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongNoteDB>(url, gongnoteDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongNoteDB>('postGongNote'))
    );
  }
  
  /** POST: add a new gongnote to the server */
  post(gongnotedb: GongNoteDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {
    return this.postGongNote(gongnotedb, GONG__StackPath, frontRepo)
  }
  postGongNote(gongnotedb: GongNoteDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    gongnotedb.GongNotePointersEncoding.Links = []
    for (let _gonglink of gongnotedb.Links) {
      gongnotedb.GongNotePointersEncoding.Links.push(_gonglink.ID)
    }
    gongnotedb.Links = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongNoteDB>(this.gongnotesUrl, gongnotedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongnotedb.Links = new Array<GongLinkDB>()
        for (let _id of gongnotedb.GongNotePointersEncoding.Links) {
          let _gonglink = frontRepo.GongLinks.get(_id)
          if (_gonglink != undefined) {
            gongnotedb.Links.push(_gonglink!)
          }
        }
        // this.log(`posted gongnotedb id=${gongnotedb.ID}`)
      }),
      catchError(this.handleError<GongNoteDB>('postGongNote'))
    );
  }

  /** DELETE: delete the gongnotedb from the server */
  delete(gongnotedb: GongNoteDB | number, GONG__StackPath: string): Observable<GongNoteDB> {
    return this.deleteGongNote(gongnotedb, GONG__StackPath)
  }
  deleteGongNote(gongnotedb: GongNoteDB | number, GONG__StackPath: string): Observable<GongNoteDB> {
    const id = typeof gongnotedb === 'number' ? gongnotedb : gongnotedb.ID;
    const url = `${this.gongnotesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongNoteDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gongnotedb id=${id}`)),
      catchError(this.handleError<GongNoteDB>('deleteGongNote'))
    );
  }

  // updateFront copy gongnote to a version with encoded pointers and update to the back
  updateFront(gongnote: GongNote, GONG__StackPath: string): Observable<GongNoteDB> {
    let gongnoteDB = new GongNoteDB
    CopyGongNoteToGongNoteDB(gongnote, gongnoteDB)
    const id = typeof gongnoteDB === 'number' ? gongnoteDB : gongnoteDB.ID
    const url = `${this.gongnotesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<GongNoteDB>(url, gongnoteDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongNoteDB>('updateGongNote'))
    );
  }

  /** PUT: update the gongnotedb on the server */
  update(gongnotedb: GongNoteDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {
    return this.updateGongNote(gongnotedb, GONG__StackPath, frontRepo)
  }
  updateGongNote(gongnotedb: GongNoteDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongNoteDB> {
    const id = typeof gongnotedb === 'number' ? gongnotedb : gongnotedb.ID;
    const url = `${this.gongnotesUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    gongnotedb.GongNotePointersEncoding.Links = []
    for (let _gonglink of gongnotedb.Links) {
      gongnotedb.GongNotePointersEncoding.Links.push(_gonglink.ID)
    }
    gongnotedb.Links = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongNoteDB>(url, gongnotedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongnotedb.Links = new Array<GongLinkDB>()
        for (let _id of gongnotedb.GongNotePointersEncoding.Links) {
          let _gonglink = frontRepo.GongLinks.get(_id)
          if (_gonglink != undefined) {
            gongnotedb.Links.push(_gonglink!)
          }
        }
        // this.log(`updated gongnotedb id=${gongnotedb.ID}`)
      }),
      catchError(this.handleError<GongNoteDB>('updateGongNote'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongNoteService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongNoteService" + error); // log to console instead

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
