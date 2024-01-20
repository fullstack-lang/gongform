import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AnimateDB } from './animate-db'
import { Animate, CopyAnimateDBToAnimate } from './animate'
import { AnimateService } from './animate.service'

import { CircleDB } from './circle-db'
import { Circle, CopyCircleDBToCircle } from './circle'
import { CircleService } from './circle.service'

import { EllipseDB } from './ellipse-db'
import { Ellipse, CopyEllipseDBToEllipse } from './ellipse'
import { EllipseService } from './ellipse.service'

import { LayerDB } from './layer-db'
import { Layer, CopyLayerDBToLayer } from './layer'
import { LayerService } from './layer.service'

import { LineDB } from './line-db'
import { Line, CopyLineDBToLine } from './line'
import { LineService } from './line.service'

import { LinkDB } from './link-db'
import { Link, CopyLinkDBToLink } from './link'
import { LinkService } from './link.service'

import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { LinkAnchoredText, CopyLinkAnchoredTextDBToLinkAnchoredText } from './linkanchoredtext'
import { LinkAnchoredTextService } from './linkanchoredtext.service'

import { PathDB } from './path-db'
import { Path, CopyPathDBToPath } from './path'
import { PathService } from './path.service'

import { PointDB } from './point-db'
import { Point, CopyPointDBToPoint } from './point'
import { PointService } from './point.service'

import { PolygoneDB } from './polygone-db'
import { Polygone, CopyPolygoneDBToPolygone } from './polygone'
import { PolygoneService } from './polygone.service'

import { PolylineDB } from './polyline-db'
import { Polyline, CopyPolylineDBToPolyline } from './polyline'
import { PolylineService } from './polyline.service'

import { RectDB } from './rect-db'
import { Rect, CopyRectDBToRect } from './rect'
import { RectService } from './rect.service'

import { RectAnchoredPathDB } from './rectanchoredpath-db'
import { RectAnchoredPath, CopyRectAnchoredPathDBToRectAnchoredPath } from './rectanchoredpath'
import { RectAnchoredPathService } from './rectanchoredpath.service'

import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { RectAnchoredRect, CopyRectAnchoredRectDBToRectAnchoredRect } from './rectanchoredrect'
import { RectAnchoredRectService } from './rectanchoredrect.service'

import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { RectAnchoredText, CopyRectAnchoredTextDBToRectAnchoredText } from './rectanchoredtext'
import { RectAnchoredTextService } from './rectanchoredtext.service'

import { RectLinkLinkDB } from './rectlinklink-db'
import { RectLinkLink, CopyRectLinkLinkDBToRectLinkLink } from './rectlinklink'
import { RectLinkLinkService } from './rectlinklink.service'

import { SVGDB } from './svg-db'
import { SVG, CopySVGDBToSVG } from './svg'
import { SVGService } from './svg.service'

import { TextDB } from './text-db'
import { Text, CopyTextDBToText } from './text'
import { TextService } from './text.service'

export const StackType = "github.com/fullstack-lang/gongsvg/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Animates_array = new Array<AnimateDB>() // array of repo instances
  Animates = new Map<number, AnimateDB>() // map of repo instances
  Animates_batch = new Map<number, AnimateDB>() // same but only in last GET (for finding repo instances to delete)

  array_Animates = new Array<Animate>() // array of front instances
  map_ID_Animate = new Map<number, Animate>() // map of front instances

  Circles_array = new Array<CircleDB>() // array of repo instances
  Circles = new Map<number, CircleDB>() // map of repo instances
  Circles_batch = new Map<number, CircleDB>() // same but only in last GET (for finding repo instances to delete)

  array_Circles = new Array<Circle>() // array of front instances
  map_ID_Circle = new Map<number, Circle>() // map of front instances

  Ellipses_array = new Array<EllipseDB>() // array of repo instances
  Ellipses = new Map<number, EllipseDB>() // map of repo instances
  Ellipses_batch = new Map<number, EllipseDB>() // same but only in last GET (for finding repo instances to delete)

  array_Ellipses = new Array<Ellipse>() // array of front instances
  map_ID_Ellipse = new Map<number, Ellipse>() // map of front instances

  Layers_array = new Array<LayerDB>() // array of repo instances
  Layers = new Map<number, LayerDB>() // map of repo instances
  Layers_batch = new Map<number, LayerDB>() // same but only in last GET (for finding repo instances to delete)

  array_Layers = new Array<Layer>() // array of front instances
  map_ID_Layer = new Map<number, Layer>() // map of front instances

  Lines_array = new Array<LineDB>() // array of repo instances
  Lines = new Map<number, LineDB>() // map of repo instances
  Lines_batch = new Map<number, LineDB>() // same but only in last GET (for finding repo instances to delete)

  array_Lines = new Array<Line>() // array of front instances
  map_ID_Line = new Map<number, Line>() // map of front instances

  Links_array = new Array<LinkDB>() // array of repo instances
  Links = new Map<number, LinkDB>() // map of repo instances
  Links_batch = new Map<number, LinkDB>() // same but only in last GET (for finding repo instances to delete)

  array_Links = new Array<Link>() // array of front instances
  map_ID_Link = new Map<number, Link>() // map of front instances

  LinkAnchoredTexts_array = new Array<LinkAnchoredTextDB>() // array of repo instances
  LinkAnchoredTexts = new Map<number, LinkAnchoredTextDB>() // map of repo instances
  LinkAnchoredTexts_batch = new Map<number, LinkAnchoredTextDB>() // same but only in last GET (for finding repo instances to delete)

  array_LinkAnchoredTexts = new Array<LinkAnchoredText>() // array of front instances
  map_ID_LinkAnchoredText = new Map<number, LinkAnchoredText>() // map of front instances

  Paths_array = new Array<PathDB>() // array of repo instances
  Paths = new Map<number, PathDB>() // map of repo instances
  Paths_batch = new Map<number, PathDB>() // same but only in last GET (for finding repo instances to delete)

  array_Paths = new Array<Path>() // array of front instances
  map_ID_Path = new Map<number, Path>() // map of front instances

  Points_array = new Array<PointDB>() // array of repo instances
  Points = new Map<number, PointDB>() // map of repo instances
  Points_batch = new Map<number, PointDB>() // same but only in last GET (for finding repo instances to delete)

  array_Points = new Array<Point>() // array of front instances
  map_ID_Point = new Map<number, Point>() // map of front instances

  Polygones_array = new Array<PolygoneDB>() // array of repo instances
  Polygones = new Map<number, PolygoneDB>() // map of repo instances
  Polygones_batch = new Map<number, PolygoneDB>() // same but only in last GET (for finding repo instances to delete)

  array_Polygones = new Array<Polygone>() // array of front instances
  map_ID_Polygone = new Map<number, Polygone>() // map of front instances

  Polylines_array = new Array<PolylineDB>() // array of repo instances
  Polylines = new Map<number, PolylineDB>() // map of repo instances
  Polylines_batch = new Map<number, PolylineDB>() // same but only in last GET (for finding repo instances to delete)

  array_Polylines = new Array<Polyline>() // array of front instances
  map_ID_Polyline = new Map<number, Polyline>() // map of front instances

  Rects_array = new Array<RectDB>() // array of repo instances
  Rects = new Map<number, RectDB>() // map of repo instances
  Rects_batch = new Map<number, RectDB>() // same but only in last GET (for finding repo instances to delete)

  array_Rects = new Array<Rect>() // array of front instances
  map_ID_Rect = new Map<number, Rect>() // map of front instances

  RectAnchoredPaths_array = new Array<RectAnchoredPathDB>() // array of repo instances
  RectAnchoredPaths = new Map<number, RectAnchoredPathDB>() // map of repo instances
  RectAnchoredPaths_batch = new Map<number, RectAnchoredPathDB>() // same but only in last GET (for finding repo instances to delete)

  array_RectAnchoredPaths = new Array<RectAnchoredPath>() // array of front instances
  map_ID_RectAnchoredPath = new Map<number, RectAnchoredPath>() // map of front instances

  RectAnchoredRects_array = new Array<RectAnchoredRectDB>() // array of repo instances
  RectAnchoredRects = new Map<number, RectAnchoredRectDB>() // map of repo instances
  RectAnchoredRects_batch = new Map<number, RectAnchoredRectDB>() // same but only in last GET (for finding repo instances to delete)

  array_RectAnchoredRects = new Array<RectAnchoredRect>() // array of front instances
  map_ID_RectAnchoredRect = new Map<number, RectAnchoredRect>() // map of front instances

  RectAnchoredTexts_array = new Array<RectAnchoredTextDB>() // array of repo instances
  RectAnchoredTexts = new Map<number, RectAnchoredTextDB>() // map of repo instances
  RectAnchoredTexts_batch = new Map<number, RectAnchoredTextDB>() // same but only in last GET (for finding repo instances to delete)

  array_RectAnchoredTexts = new Array<RectAnchoredText>() // array of front instances
  map_ID_RectAnchoredText = new Map<number, RectAnchoredText>() // map of front instances

  RectLinkLinks_array = new Array<RectLinkLinkDB>() // array of repo instances
  RectLinkLinks = new Map<number, RectLinkLinkDB>() // map of repo instances
  RectLinkLinks_batch = new Map<number, RectLinkLinkDB>() // same but only in last GET (for finding repo instances to delete)

  array_RectLinkLinks = new Array<RectLinkLink>() // array of front instances
  map_ID_RectLinkLink = new Map<number, RectLinkLink>() // map of front instances

  SVGs_array = new Array<SVGDB>() // array of repo instances
  SVGs = new Map<number, SVGDB>() // map of repo instances
  SVGs_batch = new Map<number, SVGDB>() // same but only in last GET (for finding repo instances to delete)

  array_SVGs = new Array<SVG>() // array of front instances
  map_ID_SVG = new Map<number, SVG>() // map of front instances

  Texts_array = new Array<TextDB>() // array of repo instances
  Texts = new Map<number, TextDB>() // map of repo instances
  Texts_batch = new Map<number, TextDB>() // same but only in last GET (for finding repo instances to delete)

  array_Texts = new Array<Text>() // array of front instances
  map_ID_Text = new Map<number, Text>() // map of front instances


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) { // deprecated
      // insertion point
      case 'Animate':
        return this.Animates_array as unknown as Array<Type>
      case 'Circle':
        return this.Circles_array as unknown as Array<Type>
      case 'Ellipse':
        return this.Ellipses_array as unknown as Array<Type>
      case 'Layer':
        return this.Layers_array as unknown as Array<Type>
      case 'Line':
        return this.Lines_array as unknown as Array<Type>
      case 'Link':
        return this.Links_array as unknown as Array<Type>
      case 'LinkAnchoredText':
        return this.LinkAnchoredTexts_array as unknown as Array<Type>
      case 'Path':
        return this.Paths_array as unknown as Array<Type>
      case 'Point':
        return this.Points_array as unknown as Array<Type>
      case 'Polygone':
        return this.Polygones_array as unknown as Array<Type>
      case 'Polyline':
        return this.Polylines_array as unknown as Array<Type>
      case 'Rect':
        return this.Rects_array as unknown as Array<Type>
      case 'RectAnchoredPath':
        return this.RectAnchoredPaths_array as unknown as Array<Type>
      case 'RectAnchoredRect':
        return this.RectAnchoredRects_array as unknown as Array<Type>
      case 'RectAnchoredText':
        return this.RectAnchoredTexts_array as unknown as Array<Type>
      case 'RectLinkLink':
        return this.RectLinkLinks_array as unknown as Array<Type>
      case 'SVG':
        return this.SVGs_array as unknown as Array<Type>
      case 'Text':
        return this.Texts_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  getFrontArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Animate':
        return this.array_Animates as unknown as Array<Type>
      case 'Circle':
        return this.array_Circles as unknown as Array<Type>
      case 'Ellipse':
        return this.array_Ellipses as unknown as Array<Type>
      case 'Layer':
        return this.array_Layers as unknown as Array<Type>
      case 'Line':
        return this.array_Lines as unknown as Array<Type>
      case 'Link':
        return this.array_Links as unknown as Array<Type>
      case 'LinkAnchoredText':
        return this.array_LinkAnchoredTexts as unknown as Array<Type>
      case 'Path':
        return this.array_Paths as unknown as Array<Type>
      case 'Point':
        return this.array_Points as unknown as Array<Type>
      case 'Polygone':
        return this.array_Polygones as unknown as Array<Type>
      case 'Polyline':
        return this.array_Polylines as unknown as Array<Type>
      case 'Rect':
        return this.array_Rects as unknown as Array<Type>
      case 'RectAnchoredPath':
        return this.array_RectAnchoredPaths as unknown as Array<Type>
      case 'RectAnchoredRect':
        return this.array_RectAnchoredRects as unknown as Array<Type>
      case 'RectAnchoredText':
        return this.array_RectAnchoredTexts as unknown as Array<Type>
      case 'RectLinkLink':
        return this.array_RectLinkLinks as unknown as Array<Type>
      case 'SVG':
        return this.array_SVGs as unknown as Array<Type>
      case 'Text':
        return this.array_Texts as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> { // deprecated
    switch (gongStructName) {
      // insertion point
      case 'Animate':
        return this.Animates as unknown as Map<number, Type>
      case 'Circle':
        return this.Circles as unknown as Map<number, Type>
      case 'Ellipse':
        return this.Ellipses as unknown as Map<number, Type>
      case 'Layer':
        return this.Layers as unknown as Map<number, Type>
      case 'Line':
        return this.Lines as unknown as Map<number, Type>
      case 'Link':
        return this.Links as unknown as Map<number, Type>
      case 'LinkAnchoredText':
        return this.LinkAnchoredTexts as unknown as Map<number, Type>
      case 'Path':
        return this.Paths as unknown as Map<number, Type>
      case 'Point':
        return this.Points as unknown as Map<number, Type>
      case 'Polygone':
        return this.Polygones as unknown as Map<number, Type>
      case 'Polyline':
        return this.Polylines as unknown as Map<number, Type>
      case 'Rect':
        return this.Rects as unknown as Map<number, Type>
      case 'RectAnchoredPath':
        return this.RectAnchoredPaths as unknown as Map<number, Type>
      case 'RectAnchoredRect':
        return this.RectAnchoredRects as unknown as Map<number, Type>
      case 'RectAnchoredText':
        return this.RectAnchoredTexts as unknown as Map<number, Type>
      case 'RectLinkLink':
        return this.RectLinkLinks as unknown as Map<number, Type>
      case 'SVG':
        return this.SVGs as unknown as Map<number, Type>
      case 'Text':
        return this.Texts as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
  
  getFrontMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Animate':
        return this.map_ID_Animate as unknown as Map<number, Type>
      case 'Circle':
        return this.map_ID_Circle as unknown as Map<number, Type>
      case 'Ellipse':
        return this.map_ID_Ellipse as unknown as Map<number, Type>
      case 'Layer':
        return this.map_ID_Layer as unknown as Map<number, Type>
      case 'Line':
        return this.map_ID_Line as unknown as Map<number, Type>
      case 'Link':
        return this.map_ID_Link as unknown as Map<number, Type>
      case 'LinkAnchoredText':
        return this.map_ID_LinkAnchoredText as unknown as Map<number, Type>
      case 'Path':
        return this.map_ID_Path as unknown as Map<number, Type>
      case 'Point':
        return this.map_ID_Point as unknown as Map<number, Type>
      case 'Polygone':
        return this.map_ID_Polygone as unknown as Map<number, Type>
      case 'Polyline':
        return this.map_ID_Polyline as unknown as Map<number, Type>
      case 'Rect':
        return this.map_ID_Rect as unknown as Map<number, Type>
      case 'RectAnchoredPath':
        return this.map_ID_RectAnchoredPath as unknown as Map<number, Type>
      case 'RectAnchoredRect':
        return this.map_ID_RectAnchoredRect as unknown as Map<number, Type>
      case 'RectAnchoredText':
        return this.map_ID_RectAnchoredText as unknown as Map<number, Type>
      case 'RectLinkLink':
        return this.map_ID_RectLinkLink as unknown as Map<number, Type>
      case 'SVG':
        return this.map_ID_SVG as unknown as Map<number, Type>
      case 'Text':
        return this.map_ID_Text as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private animateService: AnimateService,
    private circleService: CircleService,
    private ellipseService: EllipseService,
    private layerService: LayerService,
    private lineService: LineService,
    private linkService: LinkService,
    private linkanchoredtextService: LinkAnchoredTextService,
    private pathService: PathService,
    private pointService: PointService,
    private polygoneService: PolygoneService,
    private polylineService: PolylineService,
    private rectService: RectService,
    private rectanchoredpathService: RectAnchoredPathService,
    private rectanchoredrectService: RectAnchoredRectService,
    private rectanchoredtextService: RectAnchoredTextService,
    private rectlinklinkService: RectLinkLinkService,
    private svgService: SVGService,
    private textService: TextService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template 
    Observable<AnimateDB[]>,
    Observable<CircleDB[]>,
    Observable<EllipseDB[]>,
    Observable<LayerDB[]>,
    Observable<LineDB[]>,
    Observable<LinkDB[]>,
    Observable<LinkAnchoredTextDB[]>,
    Observable<PathDB[]>,
    Observable<PointDB[]>,
    Observable<PolygoneDB[]>,
    Observable<PolylineDB[]>,
    Observable<RectDB[]>,
    Observable<RectAnchoredPathDB[]>,
    Observable<RectAnchoredRectDB[]>,
    Observable<RectAnchoredTextDB[]>,
    Observable<RectLinkLinkDB[]>,
    Observable<SVGDB[]>,
    Observable<TextDB[]>,
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template
      this.animateService.getAnimates(this.GONG__StackPath, this.frontRepo),
      this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
      this.ellipseService.getEllipses(this.GONG__StackPath, this.frontRepo),
      this.layerService.getLayers(this.GONG__StackPath, this.frontRepo),
      this.lineService.getLines(this.GONG__StackPath, this.frontRepo),
      this.linkService.getLinks(this.GONG__StackPath, this.frontRepo),
      this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath, this.frontRepo),
      this.pathService.getPaths(this.GONG__StackPath, this.frontRepo),
      this.pointService.getPoints(this.GONG__StackPath, this.frontRepo),
      this.polygoneService.getPolygones(this.GONG__StackPath, this.frontRepo),
      this.polylineService.getPolylines(this.GONG__StackPath, this.frontRepo),
      this.rectService.getRects(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredpathService.getRectAnchoredPaths(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath, this.frontRepo),
      this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath, this.frontRepo),
      this.svgService.getSVGs(this.GONG__StackPath, this.frontRepo),
      this.textService.getTexts(this.GONG__StackPath, this.frontRepo),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template
      this.animateService.getAnimates(this.GONG__StackPath, this.frontRepo),
      this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
      this.ellipseService.getEllipses(this.GONG__StackPath, this.frontRepo),
      this.layerService.getLayers(this.GONG__StackPath, this.frontRepo),
      this.lineService.getLines(this.GONG__StackPath, this.frontRepo),
      this.linkService.getLinks(this.GONG__StackPath, this.frontRepo),
      this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath, this.frontRepo),
      this.pathService.getPaths(this.GONG__StackPath, this.frontRepo),
      this.pointService.getPoints(this.GONG__StackPath, this.frontRepo),
      this.polygoneService.getPolygones(this.GONG__StackPath, this.frontRepo),
      this.polylineService.getPolylines(this.GONG__StackPath, this.frontRepo),
      this.rectService.getRects(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredpathService.getRectAnchoredPaths(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath, this.frontRepo),
      this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath, this.frontRepo),
      this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath, this.frontRepo),
      this.svgService.getSVGs(this.GONG__StackPath, this.frontRepo),
      this.textService.getTexts(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            animates_,
            circles_,
            ellipses_,
            layers_,
            lines_,
            links_,
            linkanchoredtexts_,
            paths_,
            points_,
            polygones_,
            polylines_,
            rects_,
            rectanchoredpaths_,
            rectanchoredrects_,
            rectanchoredtexts_,
            rectlinklinks_,
            svgs_,
            texts_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var animates: AnimateDB[]
            animates = animates_ as AnimateDB[]
            var circles: CircleDB[]
            circles = circles_ as CircleDB[]
            var ellipses: EllipseDB[]
            ellipses = ellipses_ as EllipseDB[]
            var layers: LayerDB[]
            layers = layers_ as LayerDB[]
            var lines: LineDB[]
            lines = lines_ as LineDB[]
            var links: LinkDB[]
            links = links_ as LinkDB[]
            var linkanchoredtexts: LinkAnchoredTextDB[]
            linkanchoredtexts = linkanchoredtexts_ as LinkAnchoredTextDB[]
            var paths: PathDB[]
            paths = paths_ as PathDB[]
            var points: PointDB[]
            points = points_ as PointDB[]
            var polygones: PolygoneDB[]
            polygones = polygones_ as PolygoneDB[]
            var polylines: PolylineDB[]
            polylines = polylines_ as PolylineDB[]
            var rects: RectDB[]
            rects = rects_ as RectDB[]
            var rectanchoredpaths: RectAnchoredPathDB[]
            rectanchoredpaths = rectanchoredpaths_ as RectAnchoredPathDB[]
            var rectanchoredrects: RectAnchoredRectDB[]
            rectanchoredrects = rectanchoredrects_ as RectAnchoredRectDB[]
            var rectanchoredtexts: RectAnchoredTextDB[]
            rectanchoredtexts = rectanchoredtexts_ as RectAnchoredTextDB[]
            var rectlinklinks: RectLinkLinkDB[]
            rectlinklinks = rectlinklinks_ as RectLinkLinkDB[]
            var svgs: SVGDB[]
            svgs = svgs_ as SVGDB[]
            var texts: TextDB[]
            texts = texts_ as TextDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Animates_array = animates

            // clear the map that counts Animate in the GET
            this.frontRepo.Animates_batch.clear()

            animates.forEach(
              animateDB => {
                this.frontRepo.Animates.set(animateDB.ID, animateDB)
                this.frontRepo.Animates_batch.set(animateDB.ID, animateDB)
              }
            )

            // clear animates that are absent from the batch
            this.frontRepo.Animates.forEach(
              animateDB => {
                if (this.frontRepo.Animates_batch.get(animateDB.ID) == undefined) {
                  this.frontRepo.Animates.delete(animateDB.ID)
                }
              }
            )

            // sort Animates_array array
            this.frontRepo.Animates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Circles_array = circles

            // clear the map that counts Circle in the GET
            this.frontRepo.Circles_batch.clear()

            circles.forEach(
              circleDB => {
                this.frontRepo.Circles.set(circleDB.ID, circleDB)
                this.frontRepo.Circles_batch.set(circleDB.ID, circleDB)
              }
            )

            // clear circles that are absent from the batch
            this.frontRepo.Circles.forEach(
              circleDB => {
                if (this.frontRepo.Circles_batch.get(circleDB.ID) == undefined) {
                  this.frontRepo.Circles.delete(circleDB.ID)
                }
              }
            )

            // sort Circles_array array
            this.frontRepo.Circles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            this.frontRepo.Ellipses_batch.clear()

            ellipses.forEach(
              ellipseDB => {
                this.frontRepo.Ellipses.set(ellipseDB.ID, ellipseDB)
                this.frontRepo.Ellipses_batch.set(ellipseDB.ID, ellipseDB)
              }
            )

            // clear ellipses that are absent from the batch
            this.frontRepo.Ellipses.forEach(
              ellipseDB => {
                if (this.frontRepo.Ellipses_batch.get(ellipseDB.ID) == undefined) {
                  this.frontRepo.Ellipses.delete(ellipseDB.ID)
                }
              }
            )

            // sort Ellipses_array array
            this.frontRepo.Ellipses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Layers_array = layers

            // clear the map that counts Layer in the GET
            this.frontRepo.Layers_batch.clear()

            layers.forEach(
              layerDB => {
                this.frontRepo.Layers.set(layerDB.ID, layerDB)
                this.frontRepo.Layers_batch.set(layerDB.ID, layerDB)
              }
            )

            // clear layers that are absent from the batch
            this.frontRepo.Layers.forEach(
              layerDB => {
                if (this.frontRepo.Layers_batch.get(layerDB.ID) == undefined) {
                  this.frontRepo.Layers.delete(layerDB.ID)
                }
              }
            )

            // sort Layers_array array
            this.frontRepo.Layers_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Lines_array = lines

            // clear the map that counts Line in the GET
            this.frontRepo.Lines_batch.clear()

            lines.forEach(
              lineDB => {
                this.frontRepo.Lines.set(lineDB.ID, lineDB)
                this.frontRepo.Lines_batch.set(lineDB.ID, lineDB)
              }
            )

            // clear lines that are absent from the batch
            this.frontRepo.Lines.forEach(
              lineDB => {
                if (this.frontRepo.Lines_batch.get(lineDB.ID) == undefined) {
                  this.frontRepo.Lines.delete(lineDB.ID)
                }
              }
            )

            // sort Lines_array array
            this.frontRepo.Lines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            links.forEach(
              linkDB => {
                this.frontRepo.Links.set(linkDB.ID, linkDB)
                this.frontRepo.Links_batch.set(linkDB.ID, linkDB)
              }
            )

            // clear links that are absent from the batch
            this.frontRepo.Links.forEach(
              linkDB => {
                if (this.frontRepo.Links_batch.get(linkDB.ID) == undefined) {
                  this.frontRepo.Links.delete(linkDB.ID)
                }
              }
            )

            // sort Links_array array
            this.frontRepo.Links_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.LinkAnchoredTexts_array = linkanchoredtexts

            // clear the map that counts LinkAnchoredText in the GET
            this.frontRepo.LinkAnchoredTexts_batch.clear()

            linkanchoredtexts.forEach(
              linkanchoredtextDB => {
                this.frontRepo.LinkAnchoredTexts.set(linkanchoredtextDB.ID, linkanchoredtextDB)
                this.frontRepo.LinkAnchoredTexts_batch.set(linkanchoredtextDB.ID, linkanchoredtextDB)
              }
            )

            // clear linkanchoredtexts that are absent from the batch
            this.frontRepo.LinkAnchoredTexts.forEach(
              linkanchoredtextDB => {
                if (this.frontRepo.LinkAnchoredTexts_batch.get(linkanchoredtextDB.ID) == undefined) {
                  this.frontRepo.LinkAnchoredTexts.delete(linkanchoredtextDB.ID)
                }
              }
            )

            // sort LinkAnchoredTexts_array array
            this.frontRepo.LinkAnchoredTexts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Paths_array = paths

            // clear the map that counts Path in the GET
            this.frontRepo.Paths_batch.clear()

            paths.forEach(
              pathDB => {
                this.frontRepo.Paths.set(pathDB.ID, pathDB)
                this.frontRepo.Paths_batch.set(pathDB.ID, pathDB)
              }
            )

            // clear paths that are absent from the batch
            this.frontRepo.Paths.forEach(
              pathDB => {
                if (this.frontRepo.Paths_batch.get(pathDB.ID) == undefined) {
                  this.frontRepo.Paths.delete(pathDB.ID)
                }
              }
            )

            // sort Paths_array array
            this.frontRepo.Paths_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Points_array = points

            // clear the map that counts Point in the GET
            this.frontRepo.Points_batch.clear()

            points.forEach(
              pointDB => {
                this.frontRepo.Points.set(pointDB.ID, pointDB)
                this.frontRepo.Points_batch.set(pointDB.ID, pointDB)
              }
            )

            // clear points that are absent from the batch
            this.frontRepo.Points.forEach(
              pointDB => {
                if (this.frontRepo.Points_batch.get(pointDB.ID) == undefined) {
                  this.frontRepo.Points.delete(pointDB.ID)
                }
              }
            )

            // sort Points_array array
            this.frontRepo.Points_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            this.frontRepo.Polygones_batch.clear()

            polygones.forEach(
              polygoneDB => {
                this.frontRepo.Polygones.set(polygoneDB.ID, polygoneDB)
                this.frontRepo.Polygones_batch.set(polygoneDB.ID, polygoneDB)
              }
            )

            // clear polygones that are absent from the batch
            this.frontRepo.Polygones.forEach(
              polygoneDB => {
                if (this.frontRepo.Polygones_batch.get(polygoneDB.ID) == undefined) {
                  this.frontRepo.Polygones.delete(polygoneDB.ID)
                }
              }
            )

            // sort Polygones_array array
            this.frontRepo.Polygones_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            this.frontRepo.Polylines_batch.clear()

            polylines.forEach(
              polylineDB => {
                this.frontRepo.Polylines.set(polylineDB.ID, polylineDB)
                this.frontRepo.Polylines_batch.set(polylineDB.ID, polylineDB)
              }
            )

            // clear polylines that are absent from the batch
            this.frontRepo.Polylines.forEach(
              polylineDB => {
                if (this.frontRepo.Polylines_batch.get(polylineDB.ID) == undefined) {
                  this.frontRepo.Polylines.delete(polylineDB.ID)
                }
              }
            )

            // sort Polylines_array array
            this.frontRepo.Polylines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Rects_array = rects

            // clear the map that counts Rect in the GET
            this.frontRepo.Rects_batch.clear()

            rects.forEach(
              rectDB => {
                this.frontRepo.Rects.set(rectDB.ID, rectDB)
                this.frontRepo.Rects_batch.set(rectDB.ID, rectDB)
              }
            )

            // clear rects that are absent from the batch
            this.frontRepo.Rects.forEach(
              rectDB => {
                if (this.frontRepo.Rects_batch.get(rectDB.ID) == undefined) {
                  this.frontRepo.Rects.delete(rectDB.ID)
                }
              }
            )

            // sort Rects_array array
            this.frontRepo.Rects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectAnchoredPaths_array = rectanchoredpaths

            // clear the map that counts RectAnchoredPath in the GET
            this.frontRepo.RectAnchoredPaths_batch.clear()

            rectanchoredpaths.forEach(
              rectanchoredpathDB => {
                this.frontRepo.RectAnchoredPaths.set(rectanchoredpathDB.ID, rectanchoredpathDB)
                this.frontRepo.RectAnchoredPaths_batch.set(rectanchoredpathDB.ID, rectanchoredpathDB)
              }
            )

            // clear rectanchoredpaths that are absent from the batch
            this.frontRepo.RectAnchoredPaths.forEach(
              rectanchoredpathDB => {
                if (this.frontRepo.RectAnchoredPaths_batch.get(rectanchoredpathDB.ID) == undefined) {
                  this.frontRepo.RectAnchoredPaths.delete(rectanchoredpathDB.ID)
                }
              }
            )

            // sort RectAnchoredPaths_array array
            this.frontRepo.RectAnchoredPaths_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectAnchoredRects_array = rectanchoredrects

            // clear the map that counts RectAnchoredRect in the GET
            this.frontRepo.RectAnchoredRects_batch.clear()

            rectanchoredrects.forEach(
              rectanchoredrectDB => {
                this.frontRepo.RectAnchoredRects.set(rectanchoredrectDB.ID, rectanchoredrectDB)
                this.frontRepo.RectAnchoredRects_batch.set(rectanchoredrectDB.ID, rectanchoredrectDB)
              }
            )

            // clear rectanchoredrects that are absent from the batch
            this.frontRepo.RectAnchoredRects.forEach(
              rectanchoredrectDB => {
                if (this.frontRepo.RectAnchoredRects_batch.get(rectanchoredrectDB.ID) == undefined) {
                  this.frontRepo.RectAnchoredRects.delete(rectanchoredrectDB.ID)
                }
              }
            )

            // sort RectAnchoredRects_array array
            this.frontRepo.RectAnchoredRects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectAnchoredTexts_array = rectanchoredtexts

            // clear the map that counts RectAnchoredText in the GET
            this.frontRepo.RectAnchoredTexts_batch.clear()

            rectanchoredtexts.forEach(
              rectanchoredtextDB => {
                this.frontRepo.RectAnchoredTexts.set(rectanchoredtextDB.ID, rectanchoredtextDB)
                this.frontRepo.RectAnchoredTexts_batch.set(rectanchoredtextDB.ID, rectanchoredtextDB)
              }
            )

            // clear rectanchoredtexts that are absent from the batch
            this.frontRepo.RectAnchoredTexts.forEach(
              rectanchoredtextDB => {
                if (this.frontRepo.RectAnchoredTexts_batch.get(rectanchoredtextDB.ID) == undefined) {
                  this.frontRepo.RectAnchoredTexts.delete(rectanchoredtextDB.ID)
                }
              }
            )

            // sort RectAnchoredTexts_array array
            this.frontRepo.RectAnchoredTexts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RectLinkLinks_array = rectlinklinks

            // clear the map that counts RectLinkLink in the GET
            this.frontRepo.RectLinkLinks_batch.clear()

            rectlinklinks.forEach(
              rectlinklinkDB => {
                this.frontRepo.RectLinkLinks.set(rectlinklinkDB.ID, rectlinklinkDB)
                this.frontRepo.RectLinkLinks_batch.set(rectlinklinkDB.ID, rectlinklinkDB)
              }
            )

            // clear rectlinklinks that are absent from the batch
            this.frontRepo.RectLinkLinks.forEach(
              rectlinklinkDB => {
                if (this.frontRepo.RectLinkLinks_batch.get(rectlinklinkDB.ID) == undefined) {
                  this.frontRepo.RectLinkLinks.delete(rectlinklinkDB.ID)
                }
              }
            )

            // sort RectLinkLinks_array array
            this.frontRepo.RectLinkLinks_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            this.frontRepo.SVGs_batch.clear()

            svgs.forEach(
              svgDB => {
                this.frontRepo.SVGs.set(svgDB.ID, svgDB)
                this.frontRepo.SVGs_batch.set(svgDB.ID, svgDB)
              }
            )

            // clear svgs that are absent from the batch
            this.frontRepo.SVGs.forEach(
              svgDB => {
                if (this.frontRepo.SVGs_batch.get(svgDB.ID) == undefined) {
                  this.frontRepo.SVGs.delete(svgDB.ID)
                }
              }
            )

            // sort SVGs_array array
            this.frontRepo.SVGs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            texts.forEach(
              textDB => {
                this.frontRepo.Texts.set(textDB.ID, textDB)
                this.frontRepo.Texts_batch.set(textDB.ID, textDB)
              }
            )

            // clear texts that are absent from the batch
            this.frontRepo.Texts.forEach(
              textDB => {
                if (this.frontRepo.Texts_batch.get(textDB.ID) == undefined) {
                  this.frontRepo.Texts.delete(textDB.ID)
                }
              }
            )

            // sort Texts_array array
            this.frontRepo.Texts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            animates.forEach(
              animate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            circles.forEach(
              circle => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                circle.Animations = new Array<AnimateDB>()
                for (let _id of circle.CirclePointersEncoding.Animations) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    circle.Animations.push(_animate!)
                  }
                }
              }
            )
            ellipses.forEach(
              ellipse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                ellipse.Animates = new Array<AnimateDB>()
                for (let _id of ellipse.EllipsePointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    ellipse.Animates.push(_animate!)
                  }
                }
              }
            )
            layers.forEach(
              layer => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                layer.Rects = new Array<RectDB>()
                for (let _id of layer.LayerPointersEncoding.Rects) {
                  let _rect = this.frontRepo.Rects.get(_id)
                  if (_rect != undefined) {
                    layer.Rects.push(_rect!)
                  }
                }
                layer.Texts = new Array<TextDB>()
                for (let _id of layer.LayerPointersEncoding.Texts) {
                  let _text = this.frontRepo.Texts.get(_id)
                  if (_text != undefined) {
                    layer.Texts.push(_text!)
                  }
                }
                layer.Circles = new Array<CircleDB>()
                for (let _id of layer.LayerPointersEncoding.Circles) {
                  let _circle = this.frontRepo.Circles.get(_id)
                  if (_circle != undefined) {
                    layer.Circles.push(_circle!)
                  }
                }
                layer.Lines = new Array<LineDB>()
                for (let _id of layer.LayerPointersEncoding.Lines) {
                  let _line = this.frontRepo.Lines.get(_id)
                  if (_line != undefined) {
                    layer.Lines.push(_line!)
                  }
                }
                layer.Ellipses = new Array<EllipseDB>()
                for (let _id of layer.LayerPointersEncoding.Ellipses) {
                  let _ellipse = this.frontRepo.Ellipses.get(_id)
                  if (_ellipse != undefined) {
                    layer.Ellipses.push(_ellipse!)
                  }
                }
                layer.Polylines = new Array<PolylineDB>()
                for (let _id of layer.LayerPointersEncoding.Polylines) {
                  let _polyline = this.frontRepo.Polylines.get(_id)
                  if (_polyline != undefined) {
                    layer.Polylines.push(_polyline!)
                  }
                }
                layer.Polygones = new Array<PolygoneDB>()
                for (let _id of layer.LayerPointersEncoding.Polygones) {
                  let _polygone = this.frontRepo.Polygones.get(_id)
                  if (_polygone != undefined) {
                    layer.Polygones.push(_polygone!)
                  }
                }
                layer.Paths = new Array<PathDB>()
                for (let _id of layer.LayerPointersEncoding.Paths) {
                  let _path = this.frontRepo.Paths.get(_id)
                  if (_path != undefined) {
                    layer.Paths.push(_path!)
                  }
                }
                layer.Links = new Array<LinkDB>()
                for (let _id of layer.LayerPointersEncoding.Links) {
                  let _link = this.frontRepo.Links.get(_id)
                  if (_link != undefined) {
                    layer.Links.push(_link!)
                  }
                }
                layer.RectLinkLinks = new Array<RectLinkLinkDB>()
                for (let _id of layer.LayerPointersEncoding.RectLinkLinks) {
                  let _rectlinklink = this.frontRepo.RectLinkLinks.get(_id)
                  if (_rectlinklink != undefined) {
                    layer.RectLinkLinks.push(_rectlinklink!)
                  }
                }
              }
            )
            lines.forEach(
              line => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                line.Animates = new Array<AnimateDB>()
                for (let _id of line.LinePointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    line.Animates.push(_animate!)
                  }
                }
              }
            )
            links.forEach(
              link => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.LinkPointersEncoding.StartID.Int64)
                  if (_rect) {
                    link.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.LinkPointersEncoding.EndID.Int64)
                  if (_rect) {
                    link.End = _rect
                  }
                }
                // insertion point for pointers decoding
                link.TextAtArrowEnd = new Array<LinkAnchoredTextDB>()
                for (let _id of link.LinkPointersEncoding.TextAtArrowEnd) {
                  let _linkanchoredtext = this.frontRepo.LinkAnchoredTexts.get(_id)
                  if (_linkanchoredtext != undefined) {
                    link.TextAtArrowEnd.push(_linkanchoredtext!)
                  }
                }
                link.TextAtArrowStart = new Array<LinkAnchoredTextDB>()
                for (let _id of link.LinkPointersEncoding.TextAtArrowStart) {
                  let _linkanchoredtext = this.frontRepo.LinkAnchoredTexts.get(_id)
                  if (_linkanchoredtext != undefined) {
                    link.TextAtArrowStart.push(_linkanchoredtext!)
                  }
                }
                link.ControlPoints = new Array<PointDB>()
                for (let _id of link.LinkPointersEncoding.ControlPoints) {
                  let _point = this.frontRepo.Points.get(_id)
                  if (_point != undefined) {
                    link.ControlPoints.push(_point!)
                  }
                }
              }
            )
            linkanchoredtexts.forEach(
              linkanchoredtext => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                linkanchoredtext.Animates = new Array<AnimateDB>()
                for (let _id of linkanchoredtext.LinkAnchoredTextPointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    linkanchoredtext.Animates.push(_animate!)
                  }
                }
              }
            )
            paths.forEach(
              path => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                path.Animates = new Array<AnimateDB>()
                for (let _id of path.PathPointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    path.Animates.push(_animate!)
                  }
                }
              }
            )
            points.forEach(
              point => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            polygones.forEach(
              polygone => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                polygone.Animates = new Array<AnimateDB>()
                for (let _id of polygone.PolygonePointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    polygone.Animates.push(_animate!)
                  }
                }
              }
            )
            polylines.forEach(
              polyline => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                polyline.Animates = new Array<AnimateDB>()
                for (let _id of polyline.PolylinePointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    polyline.Animates.push(_animate!)
                  }
                }
              }
            )
            rects.forEach(
              rect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                rect.Animations = new Array<AnimateDB>()
                for (let _id of rect.RectPointersEncoding.Animations) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    rect.Animations.push(_animate!)
                  }
                }
                rect.RectAnchoredTexts = new Array<RectAnchoredTextDB>()
                for (let _id of rect.RectPointersEncoding.RectAnchoredTexts) {
                  let _rectanchoredtext = this.frontRepo.RectAnchoredTexts.get(_id)
                  if (_rectanchoredtext != undefined) {
                    rect.RectAnchoredTexts.push(_rectanchoredtext!)
                  }
                }
                rect.RectAnchoredRects = new Array<RectAnchoredRectDB>()
                for (let _id of rect.RectPointersEncoding.RectAnchoredRects) {
                  let _rectanchoredrect = this.frontRepo.RectAnchoredRects.get(_id)
                  if (_rectanchoredrect != undefined) {
                    rect.RectAnchoredRects.push(_rectanchoredrect!)
                  }
                }
                rect.RectAnchoredPaths = new Array<RectAnchoredPathDB>()
                for (let _id of rect.RectPointersEncoding.RectAnchoredPaths) {
                  let _rectanchoredpath = this.frontRepo.RectAnchoredPaths.get(_id)
                  if (_rectanchoredpath != undefined) {
                    rect.RectAnchoredPaths.push(_rectanchoredpath!)
                  }
                }
              }
            )
            rectanchoredpaths.forEach(
              rectanchoredpath => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            rectanchoredrects.forEach(
              rectanchoredrect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            rectanchoredtexts.forEach(
              rectanchoredtext => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                rectanchoredtext.Animates = new Array<AnimateDB>()
                for (let _id of rectanchoredtext.RectAnchoredTextPointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    rectanchoredtext.Animates.push(_animate!)
                  }
                }
              }
            )
            rectlinklinks.forEach(
              rectlinklink => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectlinklink.RectLinkLinkPointersEncoding.StartID.Int64)
                  if (_rect) {
                    rectlinklink.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _link = this.frontRepo.Links.get(rectlinklink.RectLinkLinkPointersEncoding.EndID.Int64)
                  if (_link) {
                    rectlinklink.End = _link
                  }
                }
                // insertion point for pointers decoding
              }
            )
            svgs.forEach(
              svg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field StartRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.SVGPointersEncoding.StartRectID.Int64)
                  if (_rect) {
                    svg.StartRect = _rect
                  }
                }
                // insertion point for pointer field EndRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.SVGPointersEncoding.EndRectID.Int64)
                  if (_rect) {
                    svg.EndRect = _rect
                  }
                }
                // insertion point for pointers decoding
                svg.Layers = new Array<LayerDB>()
                for (let _id of svg.SVGPointersEncoding.Layers) {
                  let _layer = this.frontRepo.Layers.get(_id)
                  if (_layer != undefined) {
                    svg.Layers.push(_layer!)
                  }
                }
              }
            )
            texts.forEach(
              text => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                text.Animates = new Array<AnimateDB>()
                for (let _id of text.TextPointersEncoding.Animates) {
                  let _animate = this.frontRepo.Animates.get(_id)
                  if (_animate != undefined) {
                    text.Animates.push(_animate!)
                  }
                }
              }
            )

            // 
            // Third Step: reddeem front objects
            // insertion point sub template for redeem 
            
            // init front objects
            this.frontRepo.array_Animates = []
            this.frontRepo.map_ID_Animate.clear()
            this.frontRepo.Animates_array.forEach(
              animateDB => {
                let animate = new Animate
                CopyAnimateDBToAnimate(animateDB, animate, this.frontRepo)
                this.frontRepo.array_Animates.push(animate)
                this.frontRepo.map_ID_Animate.set(animate.ID, animate)
              }
            )

            
            // init front objects
            this.frontRepo.array_Circles = []
            this.frontRepo.map_ID_Circle.clear()
            this.frontRepo.Circles_array.forEach(
              circleDB => {
                let circle = new Circle
                CopyCircleDBToCircle(circleDB, circle, this.frontRepo)
                this.frontRepo.array_Circles.push(circle)
                this.frontRepo.map_ID_Circle.set(circle.ID, circle)
              }
            )

            
            // init front objects
            this.frontRepo.array_Ellipses = []
            this.frontRepo.map_ID_Ellipse.clear()
            this.frontRepo.Ellipses_array.forEach(
              ellipseDB => {
                let ellipse = new Ellipse
                CopyEllipseDBToEllipse(ellipseDB, ellipse, this.frontRepo)
                this.frontRepo.array_Ellipses.push(ellipse)
                this.frontRepo.map_ID_Ellipse.set(ellipse.ID, ellipse)
              }
            )

            
            // init front objects
            this.frontRepo.array_Layers = []
            this.frontRepo.map_ID_Layer.clear()
            this.frontRepo.Layers_array.forEach(
              layerDB => {
                let layer = new Layer
                CopyLayerDBToLayer(layerDB, layer, this.frontRepo)
                this.frontRepo.array_Layers.push(layer)
                this.frontRepo.map_ID_Layer.set(layer.ID, layer)
              }
            )

            
            // init front objects
            this.frontRepo.array_Lines = []
            this.frontRepo.map_ID_Line.clear()
            this.frontRepo.Lines_array.forEach(
              lineDB => {
                let line = new Line
                CopyLineDBToLine(lineDB, line, this.frontRepo)
                this.frontRepo.array_Lines.push(line)
                this.frontRepo.map_ID_Line.set(line.ID, line)
              }
            )

            
            // init front objects
            this.frontRepo.array_Links = []
            this.frontRepo.map_ID_Link.clear()
            this.frontRepo.Links_array.forEach(
              linkDB => {
                let link = new Link
                CopyLinkDBToLink(linkDB, link, this.frontRepo)
                this.frontRepo.array_Links.push(link)
                this.frontRepo.map_ID_Link.set(link.ID, link)
              }
            )

            
            // init front objects
            this.frontRepo.array_LinkAnchoredTexts = []
            this.frontRepo.map_ID_LinkAnchoredText.clear()
            this.frontRepo.LinkAnchoredTexts_array.forEach(
              linkanchoredtextDB => {
                let linkanchoredtext = new LinkAnchoredText
                CopyLinkAnchoredTextDBToLinkAnchoredText(linkanchoredtextDB, linkanchoredtext, this.frontRepo)
                this.frontRepo.array_LinkAnchoredTexts.push(linkanchoredtext)
                this.frontRepo.map_ID_LinkAnchoredText.set(linkanchoredtext.ID, linkanchoredtext)
              }
            )

            
            // init front objects
            this.frontRepo.array_Paths = []
            this.frontRepo.map_ID_Path.clear()
            this.frontRepo.Paths_array.forEach(
              pathDB => {
                let path = new Path
                CopyPathDBToPath(pathDB, path, this.frontRepo)
                this.frontRepo.array_Paths.push(path)
                this.frontRepo.map_ID_Path.set(path.ID, path)
              }
            )

            
            // init front objects
            this.frontRepo.array_Points = []
            this.frontRepo.map_ID_Point.clear()
            this.frontRepo.Points_array.forEach(
              pointDB => {
                let point = new Point
                CopyPointDBToPoint(pointDB, point, this.frontRepo)
                this.frontRepo.array_Points.push(point)
                this.frontRepo.map_ID_Point.set(point.ID, point)
              }
            )

            
            // init front objects
            this.frontRepo.array_Polygones = []
            this.frontRepo.map_ID_Polygone.clear()
            this.frontRepo.Polygones_array.forEach(
              polygoneDB => {
                let polygone = new Polygone
                CopyPolygoneDBToPolygone(polygoneDB, polygone, this.frontRepo)
                this.frontRepo.array_Polygones.push(polygone)
                this.frontRepo.map_ID_Polygone.set(polygone.ID, polygone)
              }
            )

            
            // init front objects
            this.frontRepo.array_Polylines = []
            this.frontRepo.map_ID_Polyline.clear()
            this.frontRepo.Polylines_array.forEach(
              polylineDB => {
                let polyline = new Polyline
                CopyPolylineDBToPolyline(polylineDB, polyline, this.frontRepo)
                this.frontRepo.array_Polylines.push(polyline)
                this.frontRepo.map_ID_Polyline.set(polyline.ID, polyline)
              }
            )

            
            // init front objects
            this.frontRepo.array_Rects = []
            this.frontRepo.map_ID_Rect.clear()
            this.frontRepo.Rects_array.forEach(
              rectDB => {
                let rect = new Rect
                CopyRectDBToRect(rectDB, rect, this.frontRepo)
                this.frontRepo.array_Rects.push(rect)
                this.frontRepo.map_ID_Rect.set(rect.ID, rect)
              }
            )

            
            // init front objects
            this.frontRepo.array_RectAnchoredPaths = []
            this.frontRepo.map_ID_RectAnchoredPath.clear()
            this.frontRepo.RectAnchoredPaths_array.forEach(
              rectanchoredpathDB => {
                let rectanchoredpath = new RectAnchoredPath
                CopyRectAnchoredPathDBToRectAnchoredPath(rectanchoredpathDB, rectanchoredpath, this.frontRepo)
                this.frontRepo.array_RectAnchoredPaths.push(rectanchoredpath)
                this.frontRepo.map_ID_RectAnchoredPath.set(rectanchoredpath.ID, rectanchoredpath)
              }
            )

            
            // init front objects
            this.frontRepo.array_RectAnchoredRects = []
            this.frontRepo.map_ID_RectAnchoredRect.clear()
            this.frontRepo.RectAnchoredRects_array.forEach(
              rectanchoredrectDB => {
                let rectanchoredrect = new RectAnchoredRect
                CopyRectAnchoredRectDBToRectAnchoredRect(rectanchoredrectDB, rectanchoredrect, this.frontRepo)
                this.frontRepo.array_RectAnchoredRects.push(rectanchoredrect)
                this.frontRepo.map_ID_RectAnchoredRect.set(rectanchoredrect.ID, rectanchoredrect)
              }
            )

            
            // init front objects
            this.frontRepo.array_RectAnchoredTexts = []
            this.frontRepo.map_ID_RectAnchoredText.clear()
            this.frontRepo.RectAnchoredTexts_array.forEach(
              rectanchoredtextDB => {
                let rectanchoredtext = new RectAnchoredText
                CopyRectAnchoredTextDBToRectAnchoredText(rectanchoredtextDB, rectanchoredtext, this.frontRepo)
                this.frontRepo.array_RectAnchoredTexts.push(rectanchoredtext)
                this.frontRepo.map_ID_RectAnchoredText.set(rectanchoredtext.ID, rectanchoredtext)
              }
            )

            
            // init front objects
            this.frontRepo.array_RectLinkLinks = []
            this.frontRepo.map_ID_RectLinkLink.clear()
            this.frontRepo.RectLinkLinks_array.forEach(
              rectlinklinkDB => {
                let rectlinklink = new RectLinkLink
                CopyRectLinkLinkDBToRectLinkLink(rectlinklinkDB, rectlinklink, this.frontRepo)
                this.frontRepo.array_RectLinkLinks.push(rectlinklink)
                this.frontRepo.map_ID_RectLinkLink.set(rectlinklink.ID, rectlinklink)
              }
            )

            
            // init front objects
            this.frontRepo.array_SVGs = []
            this.frontRepo.map_ID_SVG.clear()
            this.frontRepo.SVGs_array.forEach(
              svgDB => {
                let svg = new SVG
                CopySVGDBToSVG(svgDB, svg, this.frontRepo)
                this.frontRepo.array_SVGs.push(svg)
                this.frontRepo.map_ID_SVG.set(svg.ID, svg)
              }
            )

            
            // init front objects
            this.frontRepo.array_Texts = []
            this.frontRepo.map_ID_Text.clear()
            this.frontRepo.Texts_array.forEach(
              textDB => {
                let text = new Text
                CopyTextDBToText(textDB, text, this.frontRepo)
                this.frontRepo.array_Texts.push(text)
                this.frontRepo.map_ID_Text.set(text.ID, text)
              }
            )



            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // AnimatePull performs a GET on Animate of the stack and redeem association pointers 
  AnimatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.animateService.getAnimates(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            animates,
          ]) => {
            // init the array
            this.frontRepo.Animates_array = animates

            // clear the map that counts Animate in the GET
            this.frontRepo.Animates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            animates.forEach(
              animate => {
                this.frontRepo.Animates.set(animate.ID, animate)
                this.frontRepo.Animates_batch.set(animate.ID, animate)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear animates that are absent from the GET
            this.frontRepo.Animates.forEach(
              animate => {
                if (this.frontRepo.Animates_batch.get(animate.ID) == undefined) {
                  this.frontRepo.Animates.delete(animate.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // CirclePull performs a GET on Circle of the stack and redeem association pointers 
  CirclePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.circleService.getCircles(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            circles,
          ]) => {
            // init the array
            this.frontRepo.Circles_array = circles

            // clear the map that counts Circle in the GET
            this.frontRepo.Circles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            circles.forEach(
              circle => {
                this.frontRepo.Circles.set(circle.ID, circle)
                this.frontRepo.Circles_batch.set(circle.ID, circle)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear circles that are absent from the GET
            this.frontRepo.Circles.forEach(
              circle => {
                if (this.frontRepo.Circles_batch.get(circle.ID) == undefined) {
                  this.frontRepo.Circles.delete(circle.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // EllipsePull performs a GET on Ellipse of the stack and redeem association pointers 
  EllipsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ellipseService.getEllipses(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            ellipses,
          ]) => {
            // init the array
            this.frontRepo.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            this.frontRepo.Ellipses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ellipses.forEach(
              ellipse => {
                this.frontRepo.Ellipses.set(ellipse.ID, ellipse)
                this.frontRepo.Ellipses_batch.set(ellipse.ID, ellipse)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear ellipses that are absent from the GET
            this.frontRepo.Ellipses.forEach(
              ellipse => {
                if (this.frontRepo.Ellipses_batch.get(ellipse.ID) == undefined) {
                  this.frontRepo.Ellipses.delete(ellipse.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LayerPull performs a GET on Layer of the stack and redeem association pointers 
  LayerPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.layerService.getLayers(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            layers,
          ]) => {
            // init the array
            this.frontRepo.Layers_array = layers

            // clear the map that counts Layer in the GET
            this.frontRepo.Layers_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            layers.forEach(
              layer => {
                this.frontRepo.Layers.set(layer.ID, layer)
                this.frontRepo.Layers_batch.set(layer.ID, layer)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear layers that are absent from the GET
            this.frontRepo.Layers.forEach(
              layer => {
                if (this.frontRepo.Layers_batch.get(layer.ID) == undefined) {
                  this.frontRepo.Layers.delete(layer.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinePull performs a GET on Line of the stack and redeem association pointers 
  LinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.lineService.getLines(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            lines,
          ]) => {
            // init the array
            this.frontRepo.Lines_array = lines

            // clear the map that counts Line in the GET
            this.frontRepo.Lines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            lines.forEach(
              line => {
                this.frontRepo.Lines.set(line.ID, line)
                this.frontRepo.Lines_batch.set(line.ID, line)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear lines that are absent from the GET
            this.frontRepo.Lines.forEach(
              line => {
                if (this.frontRepo.Lines_batch.get(line.ID) == undefined) {
                  this.frontRepo.Lines.delete(line.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinkPull performs a GET on Link of the stack and redeem association pointers 
  LinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkService.getLinks(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            links,
          ]) => {
            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            links.forEach(
              link => {
                this.frontRepo.Links.set(link.ID, link)
                this.frontRepo.Links_batch.set(link.ID, link)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.LinkPointersEncoding.StartID.Int64)
                  if (_rect) {
                    link.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _rect = this.frontRepo.Rects.get(link.LinkPointersEncoding.EndID.Int64)
                  if (_rect) {
                    link.End = _rect
                  }
                }
              }
            )

            // clear links that are absent from the GET
            this.frontRepo.Links.forEach(
              link => {
                if (this.frontRepo.Links_batch.get(link.ID) == undefined) {
                  this.frontRepo.Links.delete(link.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinkAnchoredTextPull performs a GET on LinkAnchoredText of the stack and redeem association pointers 
  LinkAnchoredTextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkanchoredtextService.getLinkAnchoredTexts(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            linkanchoredtexts,
          ]) => {
            // init the array
            this.frontRepo.LinkAnchoredTexts_array = linkanchoredtexts

            // clear the map that counts LinkAnchoredText in the GET
            this.frontRepo.LinkAnchoredTexts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            linkanchoredtexts.forEach(
              linkanchoredtext => {
                this.frontRepo.LinkAnchoredTexts.set(linkanchoredtext.ID, linkanchoredtext)
                this.frontRepo.LinkAnchoredTexts_batch.set(linkanchoredtext.ID, linkanchoredtext)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear linkanchoredtexts that are absent from the GET
            this.frontRepo.LinkAnchoredTexts.forEach(
              linkanchoredtext => {
                if (this.frontRepo.LinkAnchoredTexts_batch.get(linkanchoredtext.ID) == undefined) {
                  this.frontRepo.LinkAnchoredTexts.delete(linkanchoredtext.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PathPull performs a GET on Path of the stack and redeem association pointers 
  PathPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pathService.getPaths(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            paths,
          ]) => {
            // init the array
            this.frontRepo.Paths_array = paths

            // clear the map that counts Path in the GET
            this.frontRepo.Paths_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paths.forEach(
              path => {
                this.frontRepo.Paths.set(path.ID, path)
                this.frontRepo.Paths_batch.set(path.ID, path)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear paths that are absent from the GET
            this.frontRepo.Paths.forEach(
              path => {
                if (this.frontRepo.Paths_batch.get(path.ID) == undefined) {
                  this.frontRepo.Paths.delete(path.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PointPull performs a GET on Point of the stack and redeem association pointers 
  PointPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pointService.getPoints(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            points,
          ]) => {
            // init the array
            this.frontRepo.Points_array = points

            // clear the map that counts Point in the GET
            this.frontRepo.Points_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            points.forEach(
              point => {
                this.frontRepo.Points.set(point.ID, point)
                this.frontRepo.Points_batch.set(point.ID, point)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear points that are absent from the GET
            this.frontRepo.Points.forEach(
              point => {
                if (this.frontRepo.Points_batch.get(point.ID) == undefined) {
                  this.frontRepo.Points.delete(point.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PolygonePull performs a GET on Polygone of the stack and redeem association pointers 
  PolygonePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polygoneService.getPolygones(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            polygones,
          ]) => {
            // init the array
            this.frontRepo.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            this.frontRepo.Polygones_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polygones.forEach(
              polygone => {
                this.frontRepo.Polygones.set(polygone.ID, polygone)
                this.frontRepo.Polygones_batch.set(polygone.ID, polygone)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear polygones that are absent from the GET
            this.frontRepo.Polygones.forEach(
              polygone => {
                if (this.frontRepo.Polygones_batch.get(polygone.ID) == undefined) {
                  this.frontRepo.Polygones.delete(polygone.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PolylinePull performs a GET on Polyline of the stack and redeem association pointers 
  PolylinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polylineService.getPolylines(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            polylines,
          ]) => {
            // init the array
            this.frontRepo.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            this.frontRepo.Polylines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polylines.forEach(
              polyline => {
                this.frontRepo.Polylines.set(polyline.ID, polyline)
                this.frontRepo.Polylines_batch.set(polyline.ID, polyline)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear polylines that are absent from the GET
            this.frontRepo.Polylines.forEach(
              polyline => {
                if (this.frontRepo.Polylines_batch.get(polyline.ID) == undefined) {
                  this.frontRepo.Polylines.delete(polyline.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectPull performs a GET on Rect of the stack and redeem association pointers 
  RectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectService.getRects(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            rects,
          ]) => {
            // init the array
            this.frontRepo.Rects_array = rects

            // clear the map that counts Rect in the GET
            this.frontRepo.Rects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rects.forEach(
              rect => {
                this.frontRepo.Rects.set(rect.ID, rect)
                this.frontRepo.Rects_batch.set(rect.ID, rect)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear rects that are absent from the GET
            this.frontRepo.Rects.forEach(
              rect => {
                if (this.frontRepo.Rects_batch.get(rect.ID) == undefined) {
                  this.frontRepo.Rects.delete(rect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectAnchoredPathPull performs a GET on RectAnchoredPath of the stack and redeem association pointers 
  RectAnchoredPathPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectanchoredpathService.getRectAnchoredPaths(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            rectanchoredpaths,
          ]) => {
            // init the array
            this.frontRepo.RectAnchoredPaths_array = rectanchoredpaths

            // clear the map that counts RectAnchoredPath in the GET
            this.frontRepo.RectAnchoredPaths_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectanchoredpaths.forEach(
              rectanchoredpath => {
                this.frontRepo.RectAnchoredPaths.set(rectanchoredpath.ID, rectanchoredpath)
                this.frontRepo.RectAnchoredPaths_batch.set(rectanchoredpath.ID, rectanchoredpath)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear rectanchoredpaths that are absent from the GET
            this.frontRepo.RectAnchoredPaths.forEach(
              rectanchoredpath => {
                if (this.frontRepo.RectAnchoredPaths_batch.get(rectanchoredpath.ID) == undefined) {
                  this.frontRepo.RectAnchoredPaths.delete(rectanchoredpath.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectAnchoredRectPull performs a GET on RectAnchoredRect of the stack and redeem association pointers 
  RectAnchoredRectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectanchoredrectService.getRectAnchoredRects(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            rectanchoredrects,
          ]) => {
            // init the array
            this.frontRepo.RectAnchoredRects_array = rectanchoredrects

            // clear the map that counts RectAnchoredRect in the GET
            this.frontRepo.RectAnchoredRects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectanchoredrects.forEach(
              rectanchoredrect => {
                this.frontRepo.RectAnchoredRects.set(rectanchoredrect.ID, rectanchoredrect)
                this.frontRepo.RectAnchoredRects_batch.set(rectanchoredrect.ID, rectanchoredrect)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear rectanchoredrects that are absent from the GET
            this.frontRepo.RectAnchoredRects.forEach(
              rectanchoredrect => {
                if (this.frontRepo.RectAnchoredRects_batch.get(rectanchoredrect.ID) == undefined) {
                  this.frontRepo.RectAnchoredRects.delete(rectanchoredrect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectAnchoredTextPull performs a GET on RectAnchoredText of the stack and redeem association pointers 
  RectAnchoredTextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectanchoredtextService.getRectAnchoredTexts(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            rectanchoredtexts,
          ]) => {
            // init the array
            this.frontRepo.RectAnchoredTexts_array = rectanchoredtexts

            // clear the map that counts RectAnchoredText in the GET
            this.frontRepo.RectAnchoredTexts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectanchoredtexts.forEach(
              rectanchoredtext => {
                this.frontRepo.RectAnchoredTexts.set(rectanchoredtext.ID, rectanchoredtext)
                this.frontRepo.RectAnchoredTexts_batch.set(rectanchoredtext.ID, rectanchoredtext)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear rectanchoredtexts that are absent from the GET
            this.frontRepo.RectAnchoredTexts.forEach(
              rectanchoredtext => {
                if (this.frontRepo.RectAnchoredTexts_batch.get(rectanchoredtext.ID) == undefined) {
                  this.frontRepo.RectAnchoredTexts.delete(rectanchoredtext.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RectLinkLinkPull performs a GET on RectLinkLink of the stack and redeem association pointers 
  RectLinkLinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectlinklinkService.getRectLinkLinks(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            rectlinklinks,
          ]) => {
            // init the array
            this.frontRepo.RectLinkLinks_array = rectlinklinks

            // clear the map that counts RectLinkLink in the GET
            this.frontRepo.RectLinkLinks_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rectlinklinks.forEach(
              rectlinklink => {
                this.frontRepo.RectLinkLinks.set(rectlinklink.ID, rectlinklink)
                this.frontRepo.RectLinkLinks_batch.set(rectlinklink.ID, rectlinklink)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Start redeeming
                {
                  let _rect = this.frontRepo.Rects.get(rectlinklink.RectLinkLinkPointersEncoding.StartID.Int64)
                  if (_rect) {
                    rectlinklink.Start = _rect
                  }
                }
                // insertion point for pointer field End redeeming
                {
                  let _link = this.frontRepo.Links.get(rectlinklink.RectLinkLinkPointersEncoding.EndID.Int64)
                  if (_link) {
                    rectlinklink.End = _link
                  }
                }
              }
            )

            // clear rectlinklinks that are absent from the GET
            this.frontRepo.RectLinkLinks.forEach(
              rectlinklink => {
                if (this.frontRepo.RectLinkLinks_batch.get(rectlinklink.ID) == undefined) {
                  this.frontRepo.RectLinkLinks.delete(rectlinklink.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // SVGPull performs a GET on SVG of the stack and redeem association pointers 
  SVGPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.svgService.getSVGs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            svgs,
          ]) => {
            // init the array
            this.frontRepo.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            this.frontRepo.SVGs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            svgs.forEach(
              svg => {
                this.frontRepo.SVGs.set(svg.ID, svg)
                this.frontRepo.SVGs_batch.set(svg.ID, svg)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field StartRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.SVGPointersEncoding.StartRectID.Int64)
                  if (_rect) {
                    svg.StartRect = _rect
                  }
                }
                // insertion point for pointer field EndRect redeeming
                {
                  let _rect = this.frontRepo.Rects.get(svg.SVGPointersEncoding.EndRectID.Int64)
                  if (_rect) {
                    svg.EndRect = _rect
                  }
                }
              }
            )

            // clear svgs that are absent from the GET
            this.frontRepo.SVGs.forEach(
              svg => {
                if (this.frontRepo.SVGs_batch.get(svg.ID) == undefined) {
                  this.frontRepo.SVGs.delete(svg.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TextPull performs a GET on Text of the stack and redeem association pointers 
  TextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.textService.getTexts(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            texts,
          ]) => {
            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            texts.forEach(
              text => {
                this.frontRepo.Texts.set(text.ID, text)
                this.frontRepo.Texts_batch.set(text.ID, text)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear texts that are absent from the GET
            this.frontRepo.Texts.forEach(
              text => {
                if (this.frontRepo.Texts_batch.get(text.ID) == undefined) {
                  this.frontRepo.Texts.delete(text.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getAnimateUniqueID(id: number): number {
  return 31 * id
}
export function getCircleUniqueID(id: number): number {
  return 37 * id
}
export function getEllipseUniqueID(id: number): number {
  return 41 * id
}
export function getLayerUniqueID(id: number): number {
  return 43 * id
}
export function getLineUniqueID(id: number): number {
  return 47 * id
}
export function getLinkUniqueID(id: number): number {
  return 53 * id
}
export function getLinkAnchoredTextUniqueID(id: number): number {
  return 59 * id
}
export function getPathUniqueID(id: number): number {
  return 61 * id
}
export function getPointUniqueID(id: number): number {
  return 67 * id
}
export function getPolygoneUniqueID(id: number): number {
  return 71 * id
}
export function getPolylineUniqueID(id: number): number {
  return 73 * id
}
export function getRectUniqueID(id: number): number {
  return 79 * id
}
export function getRectAnchoredPathUniqueID(id: number): number {
  return 83 * id
}
export function getRectAnchoredRectUniqueID(id: number): number {
  return 89 * id
}
export function getRectAnchoredTextUniqueID(id: number): number {
  return 97 * id
}
export function getRectLinkLinkUniqueID(id: number): number {
  return 101 * id
}
export function getSVGUniqueID(id: number): number {
  return 103 * id
}
export function getTextUniqueID(id: number): number {
  return 107 * id
}
