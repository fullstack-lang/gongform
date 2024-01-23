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

import { LinkDB } from './link-db'
import { Link, CopyLinkToLinkDB } from './link'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { RectDB } from './rect-db'
import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { PointDB } from './point-db'

@Injectable({
  providedIn: 'root'
})
export class LinkService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LinkServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private linksUrl: string

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
    this.linksUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/links';
  }

  /** GET links from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB[]> {
    return this.getLinks(GONG__StackPath, frontRepo)
  }
  getLinks(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<LinkDB[]>(this.linksUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<LinkDB[]>('getLinks', []))
      );
  }

  /** GET link by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {
    return this.getLink(id, GONG__StackPath, frontRepo)
  }
  getLink(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.linksUrl}/${id}`;
    return this.http.get<LinkDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched link id=${id}`)),
      catchError(this.handleError<LinkDB>(`getLink id=${id}`))
    );
  }

  // postFront copy link to a version with encoded pointers and post to the back
  postFront(link: Link, GONG__StackPath: string): Observable<LinkDB> {
    let linkDB = new LinkDB
    CopyLinkToLinkDB(link, linkDB)
    const id = typeof linkDB === 'number' ? linkDB : linkDB.ID
    const url = `${this.linksUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LinkDB>(url, linkDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LinkDB>('postLink'))
    );
  }
  
  /** POST: add a new link to the server */
  post(linkdb: LinkDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {
    return this.postLink(linkdb, GONG__StackPath, frontRepo)
  }
  postLink(linkdb: LinkDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    if (linkdb.Start != undefined) {
      linkdb.LinkPointersEncoding.StartID.Int64 = linkdb.Start.ID
      linkdb.LinkPointersEncoding.StartID.Valid = true
    }
    linkdb.Start = undefined
    if (linkdb.End != undefined) {
      linkdb.LinkPointersEncoding.EndID.Int64 = linkdb.End.ID
      linkdb.LinkPointersEncoding.EndID.Valid = true
    }
    linkdb.End = undefined
    linkdb.LinkPointersEncoding.TextAtArrowEnd = []
    for (let _linkanchoredtext of linkdb.TextAtArrowEnd) {
      linkdb.LinkPointersEncoding.TextAtArrowEnd.push(_linkanchoredtext.ID)
    }
    linkdb.TextAtArrowEnd = []
    linkdb.LinkPointersEncoding.TextAtArrowStart = []
    for (let _linkanchoredtext of linkdb.TextAtArrowStart) {
      linkdb.LinkPointersEncoding.TextAtArrowStart.push(_linkanchoredtext.ID)
    }
    linkdb.TextAtArrowStart = []
    linkdb.LinkPointersEncoding.ControlPoints = []
    for (let _point of linkdb.ControlPoints) {
      linkdb.LinkPointersEncoding.ControlPoints.push(_point.ID)
    }
    linkdb.ControlPoints = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LinkDB>(this.linksUrl, linkdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linkdb.Start = frontRepo.Rects.get(linkdb.LinkPointersEncoding.StartID.Int64)
        linkdb.End = frontRepo.Rects.get(linkdb.LinkPointersEncoding.EndID.Int64)
        linkdb.TextAtArrowEnd = new Array<LinkAnchoredTextDB>()
        for (let _id of linkdb.LinkPointersEncoding.TextAtArrowEnd) {
          let _linkanchoredtext = frontRepo.LinkAnchoredTexts.get(_id)
          if (_linkanchoredtext != undefined) {
            linkdb.TextAtArrowEnd.push(_linkanchoredtext!)
          }
        }
        linkdb.TextAtArrowStart = new Array<LinkAnchoredTextDB>()
        for (let _id of linkdb.LinkPointersEncoding.TextAtArrowStart) {
          let _linkanchoredtext = frontRepo.LinkAnchoredTexts.get(_id)
          if (_linkanchoredtext != undefined) {
            linkdb.TextAtArrowStart.push(_linkanchoredtext!)
          }
        }
        linkdb.ControlPoints = new Array<PointDB>()
        for (let _id of linkdb.LinkPointersEncoding.ControlPoints) {
          let _point = frontRepo.Points.get(_id)
          if (_point != undefined) {
            linkdb.ControlPoints.push(_point!)
          }
        }
        // this.log(`posted linkdb id=${linkdb.ID}`)
      }),
      catchError(this.handleError<LinkDB>('postLink'))
    );
  }

  /** DELETE: delete the linkdb from the server */
  delete(linkdb: LinkDB | number, GONG__StackPath: string): Observable<LinkDB> {
    return this.deleteLink(linkdb, GONG__StackPath)
  }
  deleteLink(linkdb: LinkDB | number, GONG__StackPath: string): Observable<LinkDB> {
    const id = typeof linkdb === 'number' ? linkdb : linkdb.ID;
    const url = `${this.linksUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<LinkDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted linkdb id=${id}`)),
      catchError(this.handleError<LinkDB>('deleteLink'))
    );
  }

  // updateFront copy link to a version with encoded pointers and update to the back
  updateFront(link: Link, GONG__StackPath: string): Observable<LinkDB> {
    let linkDB = new LinkDB
    CopyLinkToLinkDB(link, linkDB)
    const id = typeof linkDB === 'number' ? linkDB : linkDB.ID
    const url = `${this.linksUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<LinkDB>(url, linkDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LinkDB>('updateLink'))
    );
  }

  /** PUT: update the linkdb on the server */
  update(linkdb: LinkDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {
    return this.updateLink(linkdb, GONG__StackPath, frontRepo)
  }
  updateLink(linkdb: LinkDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LinkDB> {
    const id = typeof linkdb === 'number' ? linkdb : linkdb.ID;
    const url = `${this.linksUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    if (linkdb.Start != undefined) {
      linkdb.LinkPointersEncoding.StartID.Int64 = linkdb.Start.ID
      linkdb.LinkPointersEncoding.StartID.Valid = true
    }
    linkdb.Start = undefined
    if (linkdb.End != undefined) {
      linkdb.LinkPointersEncoding.EndID.Int64 = linkdb.End.ID
      linkdb.LinkPointersEncoding.EndID.Valid = true
    }
    linkdb.End = undefined
    linkdb.LinkPointersEncoding.TextAtArrowEnd = []
    for (let _linkanchoredtext of linkdb.TextAtArrowEnd) {
      linkdb.LinkPointersEncoding.TextAtArrowEnd.push(_linkanchoredtext.ID)
    }
    linkdb.TextAtArrowEnd = []
    linkdb.LinkPointersEncoding.TextAtArrowStart = []
    for (let _linkanchoredtext of linkdb.TextAtArrowStart) {
      linkdb.LinkPointersEncoding.TextAtArrowStart.push(_linkanchoredtext.ID)
    }
    linkdb.TextAtArrowStart = []
    linkdb.LinkPointersEncoding.ControlPoints = []
    for (let _point of linkdb.ControlPoints) {
      linkdb.LinkPointersEncoding.ControlPoints.push(_point.ID)
    }
    linkdb.ControlPoints = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<LinkDB>(url, linkdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linkdb.Start = frontRepo.Rects.get(linkdb.LinkPointersEncoding.StartID.Int64)
        linkdb.End = frontRepo.Rects.get(linkdb.LinkPointersEncoding.EndID.Int64)
        linkdb.TextAtArrowEnd = new Array<LinkAnchoredTextDB>()
        for (let _id of linkdb.LinkPointersEncoding.TextAtArrowEnd) {
          let _linkanchoredtext = frontRepo.LinkAnchoredTexts.get(_id)
          if (_linkanchoredtext != undefined) {
            linkdb.TextAtArrowEnd.push(_linkanchoredtext!)
          }
        }
        linkdb.TextAtArrowStart = new Array<LinkAnchoredTextDB>()
        for (let _id of linkdb.LinkPointersEncoding.TextAtArrowStart) {
          let _linkanchoredtext = frontRepo.LinkAnchoredTexts.get(_id)
          if (_linkanchoredtext != undefined) {
            linkdb.TextAtArrowStart.push(_linkanchoredtext!)
          }
        }
        linkdb.ControlPoints = new Array<PointDB>()
        for (let _id of linkdb.LinkPointersEncoding.ControlPoints) {
          let _point = frontRepo.Points.get(_id)
          if (_point != undefined) {
            linkdb.ControlPoints.push(_point!)
          }
        }
        // this.log(`updated linkdb id=${linkdb.ID}`)
      }),
      catchError(this.handleError<LinkDB>('updateLink'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in LinkService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("LinkService" + error); // log to console instead

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
